use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::domain::models::note::Note;

#[derive(Serialize, Deserialize, Debug)]
pub struct CreateNoteRequest {
    pub title: String,
    pub description: Option<String>,
}

#[derive(Serialize, Debug)]
pub struct NoteResponse {
    pub id: Uuid,
    pub user_id: Uuid,
    pub title: String,
    pub description: Option<String>,
    pub is_visible: Option<NaiveDateTime>,
    pub created_at: NaiveDateTime,
}

impl NoteResponse {
    pub fn from_note_entity(note: &Note) -> Self {
        NoteResponse {
            id: note.id,
            user_id: note.user_id,
            title: note.title.clone(),
            description: note.description.clone(),
            is_visible: note.visible_at,
            created_at: note.created_at,
        }
    }
}
