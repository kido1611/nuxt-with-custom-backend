from pydantic import BaseModel
from datetime import datetime

from app.models.session_model import SessionEntity


class SessionResponse(BaseModel):
    id: str
    user_id: str | None
    csrf_token: str
    ip_address: str | None
    user_agent: str | None
    expired_at: datetime
    created_at: datetime | None


def session_entity_to_response(entity: SessionEntity) -> SessionResponse:
    return SessionResponse(
        id=entity.id,
        user_id=entity.user_id,
        csrf_token=entity.csrf_token,
        expired_at=entity.expired_at,
        ip_address=entity.ip_address,
        user_agent=entity.user_agent,
        created_at=entity.created_at,
    )
