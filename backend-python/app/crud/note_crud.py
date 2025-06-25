from typing import Sequence

from sqlmodel import Session, desc, select

from app.models.note_model import NoteEntity


def list_notes(db: Session, user_id: str) -> Sequence[NoteEntity]:
    select_statement = (
        select(NoteEntity)
        .where(NoteEntity.user_id == user_id)
        .order_by(desc(NoteEntity.created_at))
    )
    return db.exec(select_statement).all()


def create_note(db: Session, note_entity: NoteEntity) -> NoteEntity:
    db.add(note_entity)
    db.commit()
    db.refresh(note_entity)

    return note_entity


def delete_note(db: Session, user_id: str, note_id: str):
    get_note_statement = (
        select(NoteEntity)
        .where(NoteEntity.user_id == user_id)
        .where(NoteEntity.id == note_id)
    )
    note = db.exec(get_note_statement).first()

    if note is None:
        return

    db.delete(note)
    db.commit()
