use axum::{
    body::Body,
    extract::{Request, State},
    http::{
        HeaderMap, HeaderValue, Method, Response, StatusCode,
        header::{
            ACCESS_CONTROL_ALLOW_CREDENTIALS, ACCESS_CONTROL_ALLOW_HEADERS,
            ACCESS_CONTROL_ALLOW_METHODS, ACCESS_CONTROL_ALLOW_ORIGIN,
            ACCESS_CONTROL_REQUEST_METHOD,
        },
    },
    middleware::Next,
    response::IntoResponse,
};

use crate::{
    api::{router::ApiState, utils::url::parse_url_to_origin},
    errors::Error,
    infrastructure::config::ConfigCors,
};

pub async fn cors_middleware(
    State(state): State<ApiState>,
    req: Request,
    next: Next,
) -> Result<Response<Body>, Error> {
    let is_preflight_request = is_preflight_request(&req);

    let origin_header = req
        .headers()
        .get("Origin")
        .map(|origin| parse_url_to_origin(origin.to_str().unwrap().to_string()));

    if is_preflight_request {
        let mut preflight_response = StatusCode::NO_CONTENT.into_response();

        if !verify_origin(
            &state.config.cors,
            origin_header,
            preflight_response.headers_mut(),
        ) {
            *preflight_response.status_mut() = StatusCode::FORBIDDEN;

            return Ok(preflight_response);
        }

        if state.config.cors.allow_credentials {
            preflight_response.headers_mut().insert(
                ACCESS_CONTROL_ALLOW_CREDENTIALS,
                HeaderValue::from_str("true").unwrap(),
            );
        }

        preflight_response.headers_mut().insert(
            ACCESS_CONTROL_ALLOW_HEADERS,
            HeaderValue::from_str(state.config.cors.headers.join(", ").as_str()).unwrap(),
        );
        preflight_response.headers_mut().insert(
            ACCESS_CONTROL_ALLOW_METHODS,
            HeaderValue::from_str(state.config.cors.methods.join(", ").as_str()).unwrap(),
        );

        return Ok(preflight_response);
    }

    let mut res = next.run(req).await;

    if verify_origin(&state.config.cors, origin_header, res.headers_mut())
        && state.config.cors.allow_credentials
    {
        res.headers_mut().insert(
            ACCESS_CONTROL_ALLOW_CREDENTIALS,
            HeaderValue::from_str("true").unwrap(),
        );
    }

    Ok(res)
}

fn verify_origin(config: &ConfigCors, origin: Option<String>, headers: &mut HeaderMap) -> bool {
    headers.insert("Vary", HeaderValue::from_str("Origin").unwrap());

    if let Some(origin) = origin {
        if origin.is_empty() {
            return false;
        }

        if !config.origins.contains(&origin) {
            return false;
        }

        headers.insert(
            ACCESS_CONTROL_ALLOW_ORIGIN,
            HeaderValue::from_str(&origin).unwrap(),
        );
    } else {
        return false;
    }

    true
}

fn is_preflight_request(request: &Request) -> bool {
    let method = request.method();
    let access_control_request_method = match request.headers().get(ACCESS_CONTROL_REQUEST_METHOD) {
        Some(val) => val.to_str().unwrap(),
        None => return false,
    };

    method == Method::OPTIONS && !access_control_request_method.is_empty()
}
