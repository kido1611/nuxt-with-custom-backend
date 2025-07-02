use std::sync::Arc;

use axum::serve;
use backend_rust::{
    api::{app_state::AppState, router::create_router},
    application::use_cases::{note::NoteUseCase, session::SessionUseCase, user::UserUseCase},
    infrastructure::{
        config::AppConfig,
        database::repository::{
            note::SqliteNoteRepository, session::SqliteSessionRepository,
            user::SqliteUserRepository,
        },
    },
};
use log::info;
use sqlx::SqlitePool;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() -> Result<(), anyhow::Error> {
    env_logger::init();

    let config = AppConfig::new()?;
    let app_address = format!("0.0.0.0:{}", config.app.port);

    let pool = SqlitePool::connect(&config.database.url).await?;

    let user_repository = Arc::new(SqliteUserRepository { pool: pool.clone() });
    let note_repository = Arc::new(SqliteNoteRepository { pool: pool.clone() });
    let session_repository = Arc::new(SqliteSessionRepository { pool });

    let user_usecase = Arc::new(UserUseCase { user_repository });
    let session_usecase = Arc::new(SessionUseCase { session_repository });
    let note_usecase = Arc::new(NoteUseCase { note_repository });

    let app_state = AppState {
        config: Arc::new(config),
        user_usecase,
        session_usecase,
        note_usecase,
    };
    let router = create_router(app_state);

    let listener = TcpListener::bind(&app_address).await?;

    info!("Application started on {}", app_address);
    serve(listener, router.into_make_service()).await?;

    Ok(())
}
