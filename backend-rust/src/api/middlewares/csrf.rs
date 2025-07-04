use tower_cookies::Cookies;

use crate::{
    api::{
        middlewares::session::SessionExtension,
        router::ApiState,
        utils::{
            cookie::{create_clear_cookie, create_session_cookie},
            date::naivedatetime_to_expiration,
        },
    },
    errors::Error,
};

use axum::{
    Extension,
    body::Body,
    extract::{Request, State},
    http::Response,
    middleware::Next,
};

pub async fn csrf_middleware(
    State(state): State<ApiState>,
    Extension(session_extension): Extension<SessionExtension>,
    cookies: Cookies,
    req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    let is_method_safe: bool = req.method().is_safe();

    let csrf_token_header = req.headers().get("X-XSRF-TOKEN");

    let mut old_session_id: Option<String> = None;

    if !is_method_safe {
        let csrf_token_header = match csrf_token_header {
            Some(val) => val,
            None => {
                return Err(Error::CsrfTokenMissmatch(anyhow::anyhow!(
                    "csrf token is missing from request header"
                )));
            }
        };

        if let Some(ref session_response) = *session_extension.lock().await {
            old_session_id = Some(session_response.id.clone());

            if csrf_token_header.to_str().unwrap() != session_response.csrf_token {
                return Err(Error::CsrfTokenMissmatch(anyhow::anyhow!(
                    "session is missing"
                )));
            }
        } else {
            return Err(Error::CsrfTokenMissmatch(anyhow::anyhow!(
                "session is missing"
            )));
        }
    }

    let response = next.run(req).await;

    if let Some(ref session_response) = *session_extension.lock().await {
        cookies.add(create_session_cookie(
            &state.config,
            "XSRF-TOKEN".to_string(),
            session_response.csrf_token.clone(),
            naivedatetime_to_expiration(&session_response.expired_at),
            false,
        ));
    } else if old_session_id.is_some() {
        cookies.add(create_clear_cookie(
            &state.config,
            "XSRF-TOKEN".to_string(),
            false,
        ));
    }

    Ok(response)
}
