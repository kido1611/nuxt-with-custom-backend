from sqlmodel import Relationship, SQLModel, Field as SQLField
from datetime import datetime, timezone
from typing import TYPE_CHECKING, Optional

if TYPE_CHECKING:
    from .user_model import UserEntity


class NoteEntity(SQLModel, table=True):
    __tablename__ = "notes"  # type: ignore

    id: str = SQLField(primary_key=True)
    user_id: str = SQLField(foreign_key="users.id", ondelete="CASCADE")
    title: str
    description: str | None
    visible_at: datetime | None
    created_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc),
    )
    updated_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc)
    )
    deleted_at: datetime | None

    user: Optional["UserEntity"] = Relationship(back_populates="notes")
