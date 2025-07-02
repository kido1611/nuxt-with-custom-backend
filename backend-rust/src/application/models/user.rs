use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

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

#[derive(Deserialize, Debug)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}

#[derive(Deserialize, Debug)]
pub struct RegisterRequest {
    pub name: String,
    pub email: String,
    pub password: String,
}
