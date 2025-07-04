use config::{Config, File};
use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct AppConfig {
    pub app: ConfigApp,
    pub cors: ConfigCors,
    pub database: ConfigDatabase,
    pub session: ConfigSession,
}

#[derive(Deserialize, Debug)]
pub struct ConfigApp {
    pub port: i32,
}

#[derive(Deserialize, Debug)]
pub struct ConfigDatabase {
    pub url: String,
}

#[derive(Deserialize, Debug)]
pub struct ConfigSession {
    pub domain: String,
    pub lifetime: i64,
    pub secure: bool,
}

#[derive(Deserialize, Debug)]
pub struct ConfigCors {
    pub origins: Vec<String>,
    pub headers: Vec<String>,
    pub methods: Vec<String>,
    pub allow_credentials: bool,
}

impl AppConfig {
    pub fn new() -> Result<Self, config::ConfigError> {
        let config = Config::builder()
            .add_source(File::with_name("config"))
            .build()?;

        let app_config: AppConfig = config.try_deserialize().unwrap();

        Ok(app_config)
    }
}
