use url::Url;

pub fn parse_url_to_origin(origin: String) -> String {
    let mut parsed_url = match Url::parse(&origin) {
        Ok(val) => val,
        Err(_) => return "".to_string(),
    };

    parsed_url.set_path("/");
    parsed_url.set_query(None);

    parsed_url
        .join("")
        .unwrap()
        .to_string()
        .trim_end_matches("/")
        .to_lowercase()
}
