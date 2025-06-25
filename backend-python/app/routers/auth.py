from fastapi import APIRouter, HTTPException

from app.models.user_model import LoginRequest, RegisterRequest, UserEntity
from app.dependencies import DatabaseDep

router = APIRouter(prefix="/auth", tags=["auth"])


@router.post("/login")
async def login(request: LoginRequest, db: DatabaseDep):
    return {"message": "Login"}


@router.post("/register")
async def register(request: RegisterRequest, db: DatabaseDep):
    user = UserEntity(
        id="test-id", name=request.name, email=request.email, password=request.password
    )
    db.add(user)
    db.commit()
    return {"message": "Register"}


@router.delete("/logout")
async def logout():
    raise HTTPException(204)
