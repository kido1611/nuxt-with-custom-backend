use axum::Json;

use crate::api::models::response::ApiResponse;

pub async fn health_check() -> Json<ApiResponse<String>> {
    Json(ApiResponse::<String> {
        data: None,
        message: Some(String::from("Alive")),
    })
}
