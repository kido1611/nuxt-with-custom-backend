use async_trait::async_trait;
use sqlx::SqlitePool;
use uuid::Uuid;

use crate::domain::{
    models::note::{NewNote, Note},
    repository::note::NoteRepository,
};

pub struct SqliteNoteRepository {
    pub pool: SqlitePool,
}

#[async_trait]
impl NoteRepository for SqliteNoteRepository {
    async fn create(&self, note: NewNote) -> Result<(), sqlx::Error> {
        let id = note.id.to_string();
        let user_id = note.user_id.to_string();

        sqlx::query!(
            r#"INSERT INTO notes (id, user_id, title, description) VALUES (?,?,?,?)"#,
            id,
            user_id,
            note.title,
            note.description
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }

    async fn get(
        &self,
        user_id: String,
        note_id: String,
    ) -> Result<Option<crate::domain::models::note::Note>, sqlx::Error> {
        let row = match sqlx::query!(
            r#"
            SELECT * FROM notes WHERE id = ? AND user_id = ?;
        "#,
            note_id,
            user_id
        )
        .fetch_optional(&self.pool)
        .await?
        {
            Some(val) => val,
            None => return Ok(None),
        };

        let note = Note {
            id: Uuid::parse_str(&row.id).unwrap(),
            user_id: Uuid::parse_str(&row.user_id).unwrap(),
            title: row.title,
            description: row.description,
            visible_at: row.visible_at,
            created_at: row.created_at,
            updated_at: row.updated_at,
            deleted_at: row.deleted_at,
        };

        Ok(Some(note))
    }

    async fn list(
        &self,
        user_id: String,
    ) -> Result<Vec<crate::domain::models::note::Note>, sqlx::Error> {
        let rows = sqlx::query!("SELECT * FROM notes WHERE user_id = ?", user_id)
            .fetch_all(&self.pool)
            .await?;

        let mut notes = Vec::new();
        for row in rows {
            notes.push(Note {
                id: Uuid::parse_str(&row.id).unwrap(),
                user_id: Uuid::parse_str(&row.user_id).unwrap(),
                title: row.title,
                description: row.description,
                visible_at: row.visible_at,
                created_at: row.created_at,
                updated_at: row.updated_at,
                deleted_at: row.deleted_at,
            });
        }

        Ok(notes)
    }

    async fn delete(&self, user_id: String, note_id: String) -> Result<(), sqlx::Error> {
        sqlx::query!(
            "DELETE FROM notes where id = ? AND user_id = ?;",
            note_id,
            user_id
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }
}
