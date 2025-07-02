use async_trait::async_trait;

use crate::domain::models::user::{NewUser, User};

#[async_trait]
pub trait UserRepository: Send + Sync {
    async fn create(&self, user: NewUser) -> Result<(), sqlx::Error>;
    async fn get_by_email(&self, email: String) -> Result<Option<User>, sqlx::Error>;
    async fn get_by_id(&self, id: String) -> Result<Option<User>, sqlx::Error>;
}
