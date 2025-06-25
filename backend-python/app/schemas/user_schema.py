from datetime import datetime
from pydantic import BaseModel, Field

from app.models.user_model import UserEntity


class LoginRequest(BaseModel):
    email: str = Field(max_length=100)
    password: str = Field(max_length=100)


class RegisterRequest(LoginRequest):
    name: str = Field(max_length=100)


class UserResponse(BaseModel):
    id: str
    name: str
    email: str
    created_at: datetime | None


def user_entity_to_response(entity: UserEntity) -> UserResponse:
    return UserResponse(
        id=entity.id, name=entity.name, email=entity.email, created_at=entity.created_at
    )
