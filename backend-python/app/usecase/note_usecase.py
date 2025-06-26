from typing import List
from uuid import uuid4
from sqlmodel import Session

from app.crud.note_crud import create_note, delete_note, list_notes
from app.models.note_model import NoteEntity
from app.schemas.note_schema import NoteRequest, NoteResponse, note_entity_to_response


def list_user_notes(db: Session, user_id: str) -> List[NoteResponse]:
    notes_entity = list_notes(db, user_id)

    return [note_entity_to_response(note) for note in notes_entity]


def create_user_note(db: Session, user_id: str, request: NoteRequest) -> NoteResponse:
    data = NoteRequest.model_validate(request)

    note_entity = NoteEntity(
        id=str(uuid4()),
        user_id=user_id,
        title=data.title,
        description=data.description,
        visible_at=None,
        deleted_at=None,
    )

    note_entity = create_note(db, note_entity)
    return note_entity_to_response(note_entity)


def delete_user_note(db: Session, user_id: str, note_id: str):
    delete_note(db, user_id, note_id)
