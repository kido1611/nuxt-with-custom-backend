use axum::{Extension, Json};

use crate::{
    api::{middlewares::session::UserSessionExtension, models::response::ApiResponse},
    application::models::user::UserResponse,
    errors::Error,
};

pub async fn get_current_user(
    Extension(user_session_extension): Extension<UserSessionExtension>,
) -> Result<Json<ApiResponse<UserResponse>>, Error> {
    let unlocked_user_extension = user_session_extension.lock().await;
    if let Some(ref user_response) = *unlocked_user_extension {
        Ok(Json(ApiResponse {
            data: Some(UserResponse {
                id: user_response.id,
                name: user_response.name.clone(),
                email: user_response.email.clone(),
                created_at: user_response.created_at,
            }),
            message: None,
        }))
    } else {
        Err(Error::Unauthorized(anyhow::anyhow!(
            "user is missing from session"
        )))
    }
}
