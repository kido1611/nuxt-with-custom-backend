use async_trait::async_trait;
use sqlx::SqlitePool;
use uuid::Uuid;

use crate::domain::{models::session::Session, repository::session::SessionRepository};

pub struct SqliteSessionRepository {
    pub pool: SqlitePool,
}

#[async_trait]
impl SessionRepository for SqliteSessionRepository {
    async fn create(
        &self,
        session: crate::domain::models::session::NewSession,
    ) -> Result<(), sqlx::Error> {
        let user_id: Option<String> = session.user_id.map(|user_id| user_id.to_string());
        let _ = sqlx::query!(
            r#"INSERT INTO sessions (id, user_id, csrf_token, expired_at) VALUES (?, ?, ?, ?);"#,
            session.id,
            user_id,
            session.csrf_token,
            session.expired_at
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }

    async fn get_by_id(
        &self,
        session_id: String,
    ) -> Result<Option<crate::domain::models::session::Session>, sqlx::Error> {
        let row = match sqlx::query!(
            r#"
        SELECT * FROM sessions WHERE id = ? LIMIT 1;
        "#,
            session_id
        )
        .fetch_optional(&self.pool)
        .await?
        {
            Some(row) => row,
            None => return Ok(None),
        };

        let session = Session {
            id: row.id,
            user_id: row
                .user_id
                .map(|user_id| Uuid::parse_str(&user_id).unwrap()),
            csrf_token: row.csrf_token,
            ip_address: row.ip_address,
            user_agent: row.user_agent,
            expired_at: row.expired_at,
            created_at: row.created_at,
            updated_at: row.updated_at,
        };

        Ok(Some(session))
    }

    async fn update_expires_by_id(
        &self,
        session_id: String,
        expired_at: chrono::NaiveDateTime,
    ) -> Result<(), sqlx::Error> {
        let _ = sqlx::query!(
            r#"UPDATE sessions SET expired_at = ? WHERE id = ?"#,
            expired_at,
            session_id
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }

    async fn delete_by_id(&self, session_id: String) -> Result<(), sqlx::Error> {
        let _ = sqlx::query!("DELETE FROM sessions where id = ?", session_id)
            .execute(&self.pool)
            .await?;

        Ok(())
    }
}
