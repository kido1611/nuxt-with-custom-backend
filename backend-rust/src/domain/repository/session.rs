use async_trait::async_trait;
use chrono::NaiveDateTime;

use crate::domain::models::session::{NewSession, Session};

#[async_trait]
pub trait SessionRepository: Send + Sync {
    async fn create(&self, session: NewSession) -> Result<(), sqlx::Error>;
    async fn get_by_id(&self, session_id: String) -> Result<Option<Session>, sqlx::Error>;
    async fn update_expires_by_id(
        &self,
        session_id: String,
        expired_at: NaiveDateTime,
    ) -> Result<(), sqlx::Error>;
    async fn delete_by_id(&self, session_id: String) -> Result<(), sqlx::Error>;
}
