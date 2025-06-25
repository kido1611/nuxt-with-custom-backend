from fastapi import APIRouter, HTTPException

from app.models.note_model import NoteRequest

router = APIRouter(prefix="/notes", tags=["notes"])


@router.get("/")
async def list():
    return {"message": "List Note"}


@router.post("/")
async def create(request: NoteRequest):
    return {"message": "Create Note"}


@router.delete("/{note_id}")
async def delete(node_id: str):
    raise HTTPException(204)
