from uuid import uuid4
from fastapi import HTTPException
from sqlmodel import Session
from argon2.exceptions import VerifyMismatchError
from app.core.password_hash import get_password_hasher
from app.crud.user_crud import get_user_by_email, create_user as create_user_entity
from app.models.user_model import UserEntity
from app.schemas.user_schema import (
    LoginRequest,
    RegisterRequest,
    UserResponse,
    user_entity_to_response,
)


def create_user(db: Session, request: RegisterRequest) -> UserResponse:
    data = RegisterRequest.model_validate(request)

    existed_user = get_user_by_email(db, data.email)
    if existed_user is not None:
        raise HTTPException(status_code=409)

    password = get_password_hasher().hash(data.password)
    user = UserEntity(
        id=str(uuid4()), name=data.name, email=data.email, password=password
    )

    user = create_user_entity(db, user)
    return user_entity_to_response(user)


def check_user(db: Session, request: LoginRequest) -> UserResponse | None:
    data = LoginRequest.model_validate(request)

    existed_user = get_user_by_email(db, data.email)
    if existed_user is None:
        raise HTTPException(status_code=401)

    try:
        get_password_hasher().verify(existed_user.password, data.password)
    except VerifyMismatchError:
        raise HTTPException(status_code=401)

    return user_entity_to_response(existed_user)
