from fastapi import APIRouter, HTTPException

router = APIRouter()


@router.get("/sanctum/csrf-cookie")
def csrfToken():
    raise HTTPException(204)
