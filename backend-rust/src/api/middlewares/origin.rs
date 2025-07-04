use axum::{
    body::Body,
    extract::{Request, State},
    http::Response,
    middleware::Next,
};

use crate::{
    api::{router::ApiState, utils::url::parse_url_to_origin},
    errors::Error,
};

pub async fn origin_middleware(
    State(state): State<ApiState>,
    req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    let origin_header = match req.headers().get("Origin") {
        Some(origin) => origin,
        None => match req.headers().get("Referer") {
            Some(referer) => referer,
            None => {
                return Err(Error::Forbidden(anyhow::anyhow!(
                    "forbidden because incorrect origin"
                )));
            }
        },
    };

    let clean_origin = parse_url_to_origin(origin_header.to_str().unwrap().to_string());

    if !state.config.cors.origins.contains(&clean_origin) {
        return Err(Error::Forbidden(anyhow::anyhow!("origin forbidden")));
    }

    let res = next.run(req).await;

    Ok(res)
}
