use axum::{Extension, extract::State, http::StatusCode};

use crate::{
    api::{middlewares::session::SessionExtension, router::ApiState},
    errors::Error,
};

pub async fn sanctum_csrf(
    State(app_state): State<ApiState>,
    Extension(session_extension): Extension<SessionExtension>,
) -> Result<StatusCode, Error> {
    let mut unlocked_session_response = session_extension.lock().await;
    if unlocked_session_response.is_none() {
        let session_response = app_state
            .session_usecase
            .create_session(&app_state.config, None)
            .await?;

        *unlocked_session_response = Some(session_response);
    }

    Ok(StatusCode::NO_CONTENT)
}
