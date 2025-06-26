from datetime import datetime, timedelta
from sqlmodel import Session
from secrets import token_urlsafe

from app.core.config import Config
from app.crud.session_crud import create, delete, get, update
from app.crud.user_crud import get_by_id
from app.models.session_model import SessionEntity
from app.schemas.session_schema import SessionResponse, session_entity_to_response
from app.schemas.user_schema import UserResponse, user_entity_to_response


def delete_session(db: Session, session_id: str):
    delete(db, session_id)


def create_session(
    db: Session, config: Config, user_response: UserResponse | None
) -> SessionResponse:
    user_id = None
    if user_response is not None:
        user_id = user_response.id

    expired = datetime.now() + timedelta(minutes=config.session_lifetime)
    data = SessionEntity(
        id=token_urlsafe(32),
        csrf_token=token_urlsafe(32),
        user_id=user_id,
        expired_at=expired,
        ip_address=None,
        user_agent=None,
    )

    data = create(db, data)

    return session_entity_to_response(data)


def validate_session(
    db: Session, session_id: str
) -> tuple[SessionResponse | None, UserResponse | None]:
    session_entity = get(db, session_id)
    if session_entity is None:
        return (None, None)

    if session_entity.expired_at < datetime.now():
        return (None, None)

    session_response = session_entity_to_response(session_entity)
    if session_entity.user_id is None:
        return (session_response, None)

    user_entity = get_by_id(db, session_entity.user_id)
    if user_entity is None:
        return (session_response, None)

    user_response = user_entity_to_response(user_entity)

    return (session_response, user_response)


def extend_expired_session(
    db: Session, config: Config, session_id: str
) -> SessionResponse | None:
    session_entity = get(db, session_id)
    if session_entity is None:
        return None

    expires = datetime.now() + timedelta(minutes=config.session_lifetime)
    session_entity.expired_at = expires
    session_entity = update(db, session_entity)

    return session_entity_to_response(session_entity)
