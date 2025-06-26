from sqlmodel import Session, select
from app.models.user_model import UserEntity


def get_by_id(db: Session, id: str) -> UserEntity | None:
    return db.get(UserEntity, id)


def get_user_by_email(db: Session, email: str) -> UserEntity | None:
    user_statement = select(UserEntity).where(UserEntity.email == email)
    return db.exec(user_statement).first()


def create_user(db: Session, user_entity: UserEntity) -> UserEntity:
    db.add(user_entity)
    db.commit()
    db.refresh(user_entity)

    return user_entity
