use chrono::NaiveDateTime;
use serde::Serialize;
use uuid::Uuid;

#[derive(Serialize)]
pub struct Session {
    pub id: String,
    pub user_id: Option<Uuid>,
    pub csrf_token: String,
    pub ip_address: Option<String>,
    pub user_agent: Option<String>,
    pub expired_at: NaiveDateTime,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Serialize)]
pub struct NewSession {
    pub id: String,
    pub user_id: Option<Uuid>,
    pub csrf_token: String,
    pub expired_at: NaiveDateTime,
}
