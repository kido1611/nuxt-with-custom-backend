use chrono::NaiveDateTime;
use serde::Serialize;
use uuid::Uuid;

use crate::domain::models::session::Session;

#[derive(Serialize, Debug)]
pub struct SessionResponse {
    pub id: String,
    pub user_id: Option<Uuid>,
    pub csrf_token: String,
    pub ip_address: Option<String>,
    pub user_agent: Option<String>,
    pub expired_at: NaiveDateTime,
}

impl SessionResponse {
    pub fn from_session_entity(session: &Session) -> Self {
        SessionResponse {
            id: session.id.clone(),
            user_id: session.user_id,
            csrf_token: session.csrf_token.clone(),
            ip_address: session.ip_address.clone(),
            user_agent: session.user_agent.clone(),
            expired_at: session.expired_at,
        }
    }
}
