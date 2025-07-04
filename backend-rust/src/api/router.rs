use std::sync::Arc;

use axum::{
    Router, middleware,
    routing::{delete, get, post},
};
use tower_cookies::CookieManagerLayer;

use super::{
    app_state::AppState,
    handlers::{
        auth::{login, logout, register},
        health::health_check,
        note::{create_note, delete_note, get_note, list_notes},
        sanctum::sanctum_csrf,
        user::get_current_user,
    },
    middlewares::{
        cors::cors_middleware,
        csrf::csrf_middleware,
        logger::logger_middleware,
        origin::origin_middleware,
        session::{auth_session_middleware, guest_session_middleware, session_middleware},
    },
};

pub type ApiState = Arc<AppState>;

pub fn create_router(state: AppState) -> Router {
    let arc_state = Arc::new(state);

    let guest_router = Router::new()
        .route("/api/auth/login", post(login))
        .route("/api/auth/register", post(register))
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            guest_session_middleware,
        ));

    let auth_router = Router::new()
        .route("/user", get(get_current_user))
        .route("/auth/logout", delete(logout))
        .route("/notes", get(list_notes))
        .route("/notes", post(create_note))
        .route("/notes/{note_id}", get(get_note))
        .route("/notes/{note_id}", delete(delete_note))
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            auth_session_middleware,
        ));

    Router::new()
        .nest("/api", auth_router)
        .merge(guest_router)
        .route("/sanctum/csrf-cookie", get(sanctum_csrf))
        .route("/health", get(health_check))
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            csrf_middleware,
        ))
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            session_middleware,
        ))
        .layer(CookieManagerLayer::new())
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            origin_middleware,
        ))
        .layer(middleware::from_fn_with_state(
            arc_state.clone(),
            cors_middleware,
        ))
        .layer(middleware::from_fn(logger_middleware))
        .with_state(arc_state)
}
