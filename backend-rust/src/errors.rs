use axum::{body::Body, http::StatusCode, response::IntoResponse};

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("Database Error")]
    Database(sqlx::Error),

    #[error("Bad Request")]
    BadRequest(validator::ValidationErrors),

    #[error("Unauthorized")]
    Unauthorized(anyhow::Error),

    #[error("Forbidden")]
    Forbidden(anyhow::Error),

    #[error("Not Found")]
    NotFound(anyhow::Error),

    #[error("Conflict")]
    RegisterConflict(anyhow::Error),

    #[error("CSRF Token Missmatch")]
    CsrfTokenMissmatch(anyhow::Error),

    #[error("Unknown Error")]
    Other(anyhow::Error),
}

impl From<sqlx::Error> for Error {
    fn from(value: sqlx::Error) -> Self {
        Error::Database(value)
    }
}

impl From<validator::ValidationErrors> for Error {
    fn from(value: validator::ValidationErrors) -> Self {
        Error::BadRequest(value)
    }
}

impl IntoResponse for Error {
    fn into_response(self) -> axum::response::Response<Body> {
        match self {
            Error::Database(error) => {
                log::error!("{}", error);
                StatusCode::INTERNAL_SERVER_ERROR.into_response()
            }
            Error::Other(_) => StatusCode::INTERNAL_SERVER_ERROR.into_response(),
            Error::Unauthorized(error) => {
                log::error!("{}", error);
                StatusCode::UNAUTHORIZED.into_response()
            }
            Error::RegisterConflict(_) => StatusCode::CONFLICT.into_response(),
            Error::CsrfTokenMissmatch(error) => {
                log::error!("{}", error);
                StatusCode::from_u16(419).unwrap().into_response()
            }
            Error::Forbidden(_) => StatusCode::FORBIDDEN.into_response(),
            Error::NotFound(error) => {
                log::error!("{}", error);
                StatusCode::NOT_FOUND.into_response()
            }
            Error::BadRequest(error) => {
                log::warn!("{}", error);
                StatusCode::BAD_REQUEST.into_response()
            }
        }
    }
}
