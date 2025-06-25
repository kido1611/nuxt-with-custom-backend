from pydantic import BaseModel, Field
from datetime import datetime

from app.models.note_model import NoteEntity


class NoteRequest(BaseModel):
    title: str = Field(max_length=200)
    description: str | None = Field(max_length=2000)


class NoteResponse(BaseModel):
    id: str
    user_id: str
    title: str
    description: str | None
    is_visible: bool
    created_at: datetime | None


def note_entity_to_response(entity: NoteEntity) -> NoteResponse:
    return NoteResponse(
        id=entity.id,
        user_id=entity.user_id,
        title=entity.title,
        description=entity.description,
        is_visible=entity.visible_at is not None,
        created_at=entity.created_at,
    )
