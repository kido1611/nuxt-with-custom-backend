use chrono::NaiveDateTime;
use serde::Serialize;
use uuid::Uuid;

#[derive(Serialize, Debug)]
pub struct User {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub password: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Serialize)]
pub struct NewUser {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub password: String,
}
