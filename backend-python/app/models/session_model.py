from sqlmodel import Relationship, SQLModel, Field as SQLField
from datetime import datetime, timezone
from typing import TYPE_CHECKING, Optional

if TYPE_CHECKING:
    from .user_model import UserEntity


class SessionEntity(SQLModel, table=True):
    __tablename__ = "sessions"  # type: ignore

    id: str = SQLField(primary_key=True)
    user_id: str | None = SQLField(
        foreign_key="users.id", ondelete="CASCADE", nullable=True
    )
    csrf_token: str
    ip_address: str | None
    user_agent: str | None
    expired_at: datetime
    created_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc),
    )
    updated_at: datetime | None = SQLField(
        default_factory=lambda: datetime.now(timezone.utc)
    )

    user: Optional["UserEntity"] = Relationship(back_populates="sessions")
