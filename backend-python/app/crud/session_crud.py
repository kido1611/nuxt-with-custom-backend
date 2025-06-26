from sqlmodel import Session

from app.models.session_model import SessionEntity


def get(db: Session, session_id: str) -> SessionEntity | None:
    return db.get(SessionEntity, session_id)


def delete(db: Session, session_id: str):
    session = get(db, session_id)

    if session is None:
        return

    db.delete(session)
    db.commit()


def create(db: Session, session_entity: SessionEntity) -> SessionEntity:
    db.add(session_entity)
    db.commit()
    db.refresh(session_entity)

    return session_entity


def update(db: Session, session_entity: SessionEntity) -> SessionEntity:
    db.add(session_entity)
    db.commit()
    db.refresh(session_entity)

    return session_entity
