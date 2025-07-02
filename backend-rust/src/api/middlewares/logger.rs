use axum::{body::Body, extract::Request, http::Response, middleware::Next};
use chrono::Utc;
use log::info;

use crate::errors::Error;

pub async fn logger_middleware(req: Request, next: Next) -> Result<Response<Body>, Error> {
    let start = Utc::now();
    let path = req.uri().clone();
    let method = req.method().clone();

    let res = next.run(req).await;

    let end = Utc::now();
    let duration = (end - start).to_std().unwrap();
    let status = res.status();

    info!(
        "Method: {}\tPath: {}\tStatus: {}\tDuration: {:?}",
        method.as_str(),
        path.path(),
        status,
        duration
    );

    Ok(res)
}
