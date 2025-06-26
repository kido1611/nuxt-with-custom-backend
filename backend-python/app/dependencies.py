from typing import Annotated
from fastapi.security import APIKeyCookie

from fastapi import Depends, HTTPException, Request
from sqlmodel import Session

from app.schemas.session_schema import SessionResponse
from app.schemas.user_schema import UserResponse

from .core.config import Config, get_config
from .core.database import get_database_session

cookie_scheme = APIKeyCookie(name="app_session", auto_error=False)

SessionCookieDep = Annotated[str, Depends(cookie_scheme)]

DatabaseDep = Annotated[Session, Depends(get_database_session)]

ConfigDep = Annotated[Config, Depends(get_config)]


def get_session_state(request: Request) -> SessionResponse | None:
    try:
        state: SessionResponse | None = request.state.session
        return state
    except Exception:
        return None


def get_user_session_state(request: Request) -> UserResponse | None:
    try:
        state: UserResponse | None = request.state.session_user
        return state
    except Exception:
        return None


def get_auth_route(request: Request):
    session = get_session_state(request)
    if session is None:
        raise HTTPException(status_code=401)

    if session.user_id is None:
        raise HTTPException(status_code=401)


def get_guest_route(request: Request):
    user = get_user_session_state(request)
    if user is not None:
        raise HTTPException(status_code=403)
