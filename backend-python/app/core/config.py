from functools import lru_cache
from pydantic_settings import BaseSettings
from urllib.parse import urlparse


class Config(BaseSettings):
    cors_origins: str = "http://localhost:3000, https://google.com"
    cors_methods: str = "GET, POST, DELETE, OPTIONS"
    cors_headers: str = "Content-Type"
    cors_allow_credentials: bool = True

    database_url: str = "sqlite:///db.sqlite"

    session_domain: str = "localhost"
    session_lifetime: int = 180
    session_secure: bool = False

    def origins_list(self) -> list[str]:
        return [origin.strip() for origin in self.cors_origins.split(",")]

    def methods_list(self) -> list[str]:
        return [method.strip() for method in self.cors_methods.split(",")]

    def validate_origin(self, origin: str | None) -> bool:
        if origin is None:
            return False

        parse_origin = urlparse(origin)
        origin = f"{parse_origin.scheme}://{parse_origin.netloc}"
        return origin in self.origins_list()


config = Config()


@lru_cache
def get_config() -> Config:
    return config
