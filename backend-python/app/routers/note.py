from fastapi import APIRouter, HTTPException

from app.model.note import NoteRequest

router = APIRouter()


@router.get("/")
def list():
    return {"message": "List Note"}


@router.post("/")
def create(request: NoteRequest):
    return {"message": "Create Note"}


@router.delete("/{note_id}")
def delete(node_id: str):
    raise HTTPException(204)
