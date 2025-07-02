use std::sync::Arc;

use axum::{
    Extension,
    body::Body,
    extract::{Request, State},
    http::Response,
    middleware::Next,
};
use tokio::sync::Mutex;
use tower_cookies::Cookies;

use crate::{
    api::{
        router::ApiState,
        utils::{
            cookie::{create_clear_cookie, create_session_cookie},
            date::naivedatetime_to_expiration,
        },
    },
    application::models::{session::SessionResponse, user::UserResponse},
    errors::Error,
};

pub type SessionExtension = Arc<Mutex<Option<SessionResponse>>>;
pub type UserSessionExtension = Arc<Mutex<Option<UserResponse>>>;

pub async fn session_middleware(
    State(state): State<ApiState>,
    cookies: Cookies,
    mut req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    let session_cookie = cookies.get("app_session");

    let session_data = if let Some(session_id) = session_cookie.clone() {
        let session_response = state
            .session_usecase
            .get_session_by_id(session_id.value().to_string())
            .await?;

        if session_response.is_none() {
            // clear cookie because session cookie is exist but missing in database
            cookies.add(create_clear_cookie(
                &state.config,
                "app_session".to_string(),
                true,
            ));
        }

        session_response
    } else {
        None
    };

    let user_data: Option<UserResponse> = if let Some(session_response) = &session_data {
        if let Some(user_id) = session_response.user_id {
            state
                .user_usecase
                .get_user_by_id(user_id.to_string())
                .await?
        } else {
            None
        }
    } else {
        None
    };

    let session_extension: SessionExtension = Arc::new(Mutex::new(session_data));
    req.extensions_mut().insert(session_extension.clone());

    let user_session_extension: UserSessionExtension = Arc::new(Mutex::new(user_data));
    req.extensions_mut().insert(user_session_extension);

    let res = next.run(req).await;

    let latest_session_response = session_extension.lock().await;

    if let Some(ref session) = *latest_session_response {
        if let Some(cookie_id_value) = session_cookie {
            // different session
            if session.id != cookie_id_value.value() {
                cookies.add(create_session_cookie(
                    &state.config,
                    "app_session".to_string(),
                    session.id.clone(),
                    naivedatetime_to_expiration(&session.expired_at),
                    true,
                ));
            } else if let Some(new_session_response) = state
                .session_usecase
                .extend_session(&state.config, session)
                .await?
            {
                cookies.add(create_session_cookie(
                    &state.config,
                    "app_session".to_string(),
                    new_session_response.id.clone(),
                    naivedatetime_to_expiration(&new_session_response.expired_at),
                    true,
                ));
            }
        } else {
            // login condition
            cookies.add(create_session_cookie(
                &state.config,
                "app_session".to_string(),
                session.id.clone(),
                naivedatetime_to_expiration(&session.expired_at),
                true,
            ));
        }
    } else if session_cookie.is_some() {
        // session/user missing condition
        cookies.add(create_clear_cookie(
            &state.config,
            "app_session".to_string(),
            true,
        ));
    }

    Ok(res)
}

pub async fn guest_session_middleware(
    Extension(user_session_extension): Extension<UserSessionExtension>,
    req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    // alternative for unlocking guard without drop, wrap with scope
    // required to unlock mutex to allow using extension in next middleware/handler
    {
        let latest_user_response = user_session_extension.lock().await;

        if (*latest_user_response).is_some() {
            return Err(Error::Forbidden(anyhow::anyhow!(
                "Forbidden accessing routes from middleware"
            )));
        }
    }

    Ok(next.run(req).await)
}

pub async fn auth_session_middleware(
    Extension(user_session_extension): Extension<UserSessionExtension>,
    req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    let latest_user_response = user_session_extension.lock().await;

    if (*latest_user_response).is_none() {
        return Err(Error::Unauthorized(anyhow::anyhow!(
            "Unauthorized accessing routes from middleware"
        )));
    }

    // unlock mutex
    // required to unlock mutex to allow using extension in next middleware/handler
    drop(latest_user_response);

    Ok(next.run(req).await)
}
