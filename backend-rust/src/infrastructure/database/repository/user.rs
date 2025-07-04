use async_trait::async_trait;
use sqlx::SqlitePool;
use uuid::Uuid;

use crate::domain::{
    models::user::{NewUser, User},
    repository::user::UserRepository,
};

pub struct SqliteUserRepository {
    pub pool: SqlitePool,
}

#[async_trait]
impl UserRepository for SqliteUserRepository {
    async fn create(&self, user: NewUser) -> Result<(), sqlx::Error> {
        let id: String = user.id.to_string();
        let _ = sqlx::query!(
            r#"INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?);"#,
            id,
            user.name,
            user.email,
            user.password
        )
        .execute(&self.pool)
        .await?;

        Ok(())
    }

    async fn get_by_email(
        &self,
        email: String,
    ) -> Result<Option<crate::domain::models::user::User>, sqlx::Error> {
        let row = match sqlx::query!(
            r#"
        SELECT * FROM users WHERE email = ? LIMIT 1;
        "#,
            email
        )
        .fetch_optional(&self.pool)
        .await?
        {
            Some(row) => row,
            None => return Ok(None),
        };

        let user = User {
            id: Uuid::parse_str(&row.id).unwrap(),
            name: row.name,
            email: row.email,
            password: row.password,
            created_at: row.created_at,
            updated_at: row.updated_at,
        };

        Ok(Some(user))
    }

    async fn get_by_id(
        &self,
        id: String,
    ) -> Result<Option<crate::domain::models::user::User>, sqlx::Error> {
        let row = match sqlx::query!(
            r#"
        SELECT * FROM users WHERE id = ? LIMIT 1;
        "#,
            id
        )
        .fetch_optional(&self.pool)
        .await?
        {
            Some(row) => row,
            None => return Ok(None),
        };

        let user = User {
            id: Uuid::parse_str(&row.id).unwrap(),
            name: row.name,
            email: row.email,
            password: row.password,
            created_at: row.created_at,
            updated_at: row.updated_at,
        };

        Ok(Some(user))
    }
}
