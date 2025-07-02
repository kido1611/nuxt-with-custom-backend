use std::sync::Arc;

use axum::{Extension, Json, extract::State, http::StatusCode};

use crate::{
    api::{
        app_state::AppState, middlewares::session::SessionExtension, models::response::ApiResponse,
        router::ApiState,
    },
    application::models::user::{LoginRequest, RegisterRequest, UserResponse},
    errors::Error,
};

pub async fn login(
    State(app_state): State<Arc<AppState>>,
    Extension(session): Extension<SessionExtension>,
    Json(request): axum::extract::Json<LoginRequest>,
) -> Result<Json<ApiResponse<UserResponse>>, Error> {
    let user_response = match app_state.user_usecase.check_user(&request).await? {
        Some(user) => user,
        None => return Err(Error::Unauthorized(anyhow::anyhow!("User not found"))),
    };

    let session_response = app_state
        .session_usecase
        .create_session(&app_state.config, Some(&user_response))
        .await?;

    let mut unlock_session = session.lock().await;
    if let Some(ref session_response) = *unlock_session {
        app_state
            .session_usecase
            .delete_session(session_response.id.clone())
            .await?;
    }

    *unlock_session = Some(session_response);

    Ok(Json(ApiResponse {
        data: Some(user_response),
        message: None,
    }))
}

pub async fn register(
    State(app_state): State<Arc<AppState>>,
    Extension(session): Extension<SessionExtension>,
    Json(request): axum::extract::Json<RegisterRequest>,
) -> Result<(StatusCode, Json<ApiResponse<UserResponse>>), Error> {
    let user = app_state.user_usecase.create_user(&request).await?;

    let mut unlock_session = session.lock().await;
    if let Some(ref session_response) = *unlock_session {
        app_state
            .session_usecase
            .delete_session(session_response.id.clone())
            .await?;
    }

    *unlock_session = None;

    Ok((
        StatusCode::CREATED,
        Json(ApiResponse {
            data: Some(user),
            message: None,
        }),
    ))
}

pub async fn logout(
    State(app_state): State<ApiState>,
    Extension(session_extension): Extension<SessionExtension>,
) -> Result<StatusCode, Error> {
    let mut unlocked_session_response = session_extension.lock().await;
    if let Some(ref session_response) = *unlocked_session_response {
        app_state
            .session_usecase
            .delete_session(session_response.id.clone())
            .await?;
    }
    *unlocked_session_response = None;

    Ok(StatusCode::NO_CONTENT)
}
