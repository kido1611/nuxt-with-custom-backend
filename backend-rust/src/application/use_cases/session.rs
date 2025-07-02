use std::sync::Arc;

use chrono::{Duration, Utc};
use rand::{RngCore, rngs::ThreadRng};

use crate::{
    application::models::{session::SessionResponse, user::UserResponse},
    domain::{models::session::NewSession, repository::session::SessionRepository},
    errors::Error,
    infrastructure::config::AppConfig,
};

pub struct SessionUseCase {
    pub session_repository: Arc<dyn SessionRepository>,
}

impl SessionUseCase {
    pub async fn create_session(
        &self,
        app_config: &AppConfig,
        user: Option<&UserResponse>,
    ) -> Result<SessionResponse, Error> {
        let session_id = self.generate_random_token();
        let new_session = NewSession {
            id: session_id.clone(),
            user_id: user.map(|usr| usr.id),
            csrf_token: self.generate_random_token(),
            expired_at: (Utc::now() + Duration::minutes(app_config.session.lifetime)).naive_utc(),
        };

        // insert new session to database
        self.session_repository
            .create(new_session)
            .await
            .map_err(Error::Database)?;

        let session = self
            .session_repository
            .get_by_id(session_id)
            .await?
            .unwrap(); // use unwrap because session must be exists

        let session_response = SessionResponse::from_session_entity(&session);
        Ok(session_response)
    }

    pub async fn get_session_by_id(
        &self,
        session_id: String,
    ) -> Result<Option<SessionResponse>, Error> {
        if let Some(session) = self.session_repository.get_by_id(session_id).await? {
            let session_response = SessionResponse::from_session_entity(&session);
            Ok(Some(session_response))
        } else {
            Ok(None)
        }
    }

    pub async fn extend_session(
        &self,
        app_config: &AppConfig,
        session_response: &SessionResponse,
    ) -> Result<Option<SessionResponse>, Error> {
        let current_datetime = (Utc::now() - Duration::minutes(30)).naive_utc();
        if session_response.expired_at > current_datetime {
            return Ok(None);
        }

        let new_expired = (Utc::now() + Duration::minutes(app_config.session.lifetime)).naive_utc();
        self.session_repository
            .update_expires_by_id(session_response.id.clone(), new_expired)
            .await
            .map_err(Error::Database)?;

        let new_session_response = SessionResponse {
            id: session_response.id.clone(),
            csrf_token: session_response.csrf_token.clone(),
            ip_address: session_response.ip_address.clone(),
            user_agent: session_response.user_agent.clone(),
            user_id: session_response.user_id,
            expired_at: new_expired,
        };
        Ok(Some(new_session_response))
    }

    pub async fn delete_session(&self, session_id: String) -> Result<(), Error> {
        self.session_repository
            .delete_by_id(session_id)
            .await
            .map_err(Error::Database)
    }

    fn generate_random_token(&self) -> String {
        let mut rng = ThreadRng::default();
        let mut bytes = [0u8; 16];
        rng.fill_bytes(&mut bytes);

        hex::encode(bytes)
    }
}
