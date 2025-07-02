use time::{Duration, OffsetDateTime};
use tower_cookies::{Cookie, cookie::Expiration};

use crate::infrastructure::config::AppConfig;

pub fn create_session_cookie<'c, E>(
    config: &AppConfig,
    name: String,
    value: String,
    expires: E,
    http_only: bool,
) -> Cookie<'c>
where
    E: Into<Expiration>,
{
    Cookie::build((name, value))
        .domain(config.session.domain.clone())
        .expires(expires)
        .path("/")
        .http_only(http_only)
        .secure(config.session.secure)
        .same_site(tower_cookies::cookie::SameSite::Lax)
        .build()
}

pub fn create_clear_cookie<'c>(config: &AppConfig, name: String, http_only: bool) -> Cookie<'c> {
    create_session_cookie(
        config,
        name,
        "".to_string(),
        Expiration::from(OffsetDateTime::now_utc().saturating_sub(Duration::days(1))),
        http_only,
    )
}
