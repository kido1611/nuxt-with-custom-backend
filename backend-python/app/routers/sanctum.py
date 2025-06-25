from fastapi import APIRouter, HTTPException

router = APIRouter(tags=["auth"])


@router.get("/sanctum/csrf-cookie")
async def csrf_token():
    raise HTTPException(204)
