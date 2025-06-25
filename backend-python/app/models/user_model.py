from datetime import datetime, timezone
from sqlmodel import Relationship, SQLModel, Field as SQLField

from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from .session_model import SessionEntity
    from .note_model import NoteEntity


class UserEntity(SQLModel, table=True):
    __tablename__ = "users"  # type: ignore

    id: str = SQLField(primary_key=True)
    name: str
    email: str = SQLField(unique=True)
    password: str
    created_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc),
    )
    updated_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc)
    )

    sessions: list["SessionEntity"] = Relationship(
        back_populates="user", cascade_delete=True
    )
    notes: list["NoteEntity"] = Relationship(back_populates="user", cascade_delete=True)
