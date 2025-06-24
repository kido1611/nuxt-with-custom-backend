from fastapi import APIRouter, HTTPException

from app.model import user

router = APIRouter()


@router.post("/auth/login")
def login(request: user.LoginRequest):
    return {"message": "Login"}


@router.post("/auth/register")
def register(request: user.RegisterRequest):
    return {"message": "Register"}


@router.delete("/auth/logout")
def logout():
    raise HTTPException(204)
