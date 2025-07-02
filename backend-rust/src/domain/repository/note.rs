use async_trait::async_trait;

use crate::domain::models::note::{NewNote, Note};

#[async_trait]
pub trait NoteRepository: Send + Sync {
    async fn create(&self, note: NewNote) -> Result<(), sqlx::Error>;
    async fn get(&self, user_id: String, note_id: String) -> Result<Option<Note>, sqlx::Error>;
    async fn list(&self, user_id: String) -> Result<Vec<Note>, sqlx::Error>;
    async fn delete(&self, user_id: String, note_id: String) -> Result<(), sqlx::Error>;
}
