from datetime import datetime, timezone
from pydantic import BaseModel, Field
from sqlmodel import Relationship, SQLModel, Field as SQLField

from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from .note_model import NoteEntity
    from .session_model import SessionEntity


class LoginRequest(BaseModel):
    email: str = Field(max_length=100)
    password: str = Field(max_length=100)


class RegisterRequest(LoginRequest):
    name: str = Field(max_length=100)


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

    notes: list["NoteEntity"] = Relationship(back_populates="user", cascade_delete=True)
    sessions: list["SessionEntity"] = Relationship(
        back_populates="user", cascade_delete=True
    )


class UserResponse(BaseModel):
    id: str
    name: str
    email: str
    created_at: datetime | None
