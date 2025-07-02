use axum::{
    Extension, Json,
    extract::{Path, State},
    http::StatusCode,
};

use crate::{
    api::{
        middlewares::session::UserSessionExtension, models::response::ApiResponse, router::ApiState,
    },
    application::models::note::{CreateNoteRequest, NoteResponse},
    errors::Error,
};

pub async fn list_notes(
    State(api_state): State<ApiState>,
    Extension(user_session_extension): Extension<UserSessionExtension>,
) -> Result<Json<ApiResponse<Vec<NoteResponse>>>, Error> {
    let unlocked_user_extension = user_session_extension.lock().await;
    if let Some(ref user_response) = *unlocked_user_extension {
        let notes = api_state
            .note_usecase
            .list_notes(user_response.id.to_string())
            .await?;

        Ok(Json(ApiResponse {
            data: Some(notes),
            message: None,
        }))
    } else {
        Err(Error::Unauthorized(anyhow::anyhow!(
            "user is missing from session"
        )))
    }
}

pub async fn create_note(
    State(api_state): State<ApiState>,
    Extension(user_session_extension): Extension<UserSessionExtension>,
    Json(request): Json<CreateNoteRequest>,
) -> Result<(StatusCode, Json<ApiResponse<NoteResponse>>), Error> {
    let unlocked_user_extension = user_session_extension.lock().await;
    if let Some(ref user_response) = *unlocked_user_extension {
        let note = api_state
            .note_usecase
            .create_note(user_response.id.to_string(), request)
            .await?;

        Ok((
            StatusCode::CREATED,
            Json(ApiResponse {
                data: Some(note),
                message: None,
            }),
        ))
    } else {
        Err(Error::Unauthorized(anyhow::anyhow!(
            "user is missing from session"
        )))
    }
}

pub async fn get_note(
    State(api_state): State<ApiState>,
    Extension(user_session_extension): Extension<UserSessionExtension>,
    Path(note_id): Path<String>,
) -> Result<Json<ApiResponse<NoteResponse>>, Error> {
    let unlocked_user_extension = user_session_extension.lock().await;
    let user_response = if let Some(ref user_response) = *unlocked_user_extension {
        user_response
    } else {
        return Err(Error::Unauthorized(anyhow::anyhow!(
            "user is missing from session"
        )));
    };

    let note_response = match api_state
        .note_usecase
        .get_note(user_response.id.to_string(), note_id)
        .await?
    {
        Some(note_response) => note_response,
        None => {
            return Err(Error::NotFound(anyhow::anyhow!("Note is missing")));
        }
    };

    Ok(Json(ApiResponse {
        data: Some(note_response),
        message: None,
    }))
}

pub async fn delete_note(
    State(api_state): State<ApiState>,
    Extension(user_session_extension): Extension<UserSessionExtension>,
    Path(note_id): Path<String>,
) -> Result<StatusCode, Error> {
    let unlocked_user_extension = user_session_extension.lock().await;
    let user_response = if let Some(ref user_response) = *unlocked_user_extension {
        user_response
    } else {
        return Err(Error::Unauthorized(anyhow::anyhow!(
            "user is missing from session"
        )));
    };

    api_state
        .note_usecase
        .delete_note(user_response.id.to_string(), note_id)
        .await?;

    Ok(StatusCode::NO_CONTENT)
}
