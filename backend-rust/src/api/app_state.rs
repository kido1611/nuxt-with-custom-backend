use std::sync::Arc;

use crate::{
    application::use_cases::{note::NoteUseCase, session::SessionUseCase, user::UserUseCase},
    infrastructure::config::AppConfig,
};

pub struct AppState {
    pub config: Arc<AppConfig>,
    pub user_usecase: Arc<UserUseCase>,
    pub session_usecase: Arc<SessionUseCase>,
    pub note_usecase: Arc<NoteUseCase>,
}
