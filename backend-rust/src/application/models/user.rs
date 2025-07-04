use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;
use validator::Validate;

use crate::domain::models::user::User;

#[derive(Serialize, Debug)]
pub struct UserResponse {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub created_at: NaiveDateTime,
}

impl UserResponse {
    pub fn from_user_entity(user: &User) -> Self {
        UserResponse {
            id: user.id,
            name: user.name.clone(),
            email: user.email.clone(),
            created_at: user.created_at,
        }
    }
}

#[derive(Deserialize, Debug, Validate)]
pub struct LoginRequest {
    #[validate(email, length(max = 100))]
    pub email: String,
    #[validate(length(max = 32))]
    pub password: String,
}

#[derive(Deserialize, Debug, Validate)]
pub struct RegisterRequest {
    #[validate(length(max = 100))]
    pub name: String,
    #[validate(email, length(max = 100))]
    pub email: String,
    #[validate(length(max = 32))]
    pub password: String,
}
