from typing import List
from fastapi import APIRouter, Depends, HTTPException, Request
from fastapi.responses import PlainTextResponse

from app.dependencies import (
    DatabaseDep,
    SessionCookieDep,
    get_auth_route,
    get_user_session_state,
)
from app.schemas.note_schema import NoteRequest, NoteResponse
from app.schemas.response_schema import ApiResponse
from app.usecase.note_usecase import create_user_note, delete_user_note, list_user_notes

router = APIRouter(
    prefix="/notes", tags=["notes"], dependencies=[Depends(get_auth_route)]
)


@router.get("", status_code=200, response_model=ApiResponse[List[NoteResponse]])
@router.get(
    "/",
    include_in_schema=False,
    status_code=200,
    response_model=ApiResponse[List[NoteResponse]],
)
async def list(request: Request, _: SessionCookieDep, db: DatabaseDep):
    user = get_user_session_state(request)

    if user is None:
        raise HTTPException(status_code=401)

    notes = list_user_notes(db, user.id)
    return ApiResponse(data=notes)


@router.post("", status_code=201, response_model=ApiResponse[NoteResponse])
@router.post(
    "/",
    include_in_schema=False,
    status_code=201,
    response_model=ApiResponse[NoteResponse],
)
async def create(
    note_request: NoteRequest, request: Request, _: SessionCookieDep, db: DatabaseDep
):
    user = get_user_session_state(request)

    if user is None:
        raise HTTPException(status_code=401)

    note_response = create_user_note(db, user.id, note_request)
    return ApiResponse(data=note_response)


@router.delete("/{note_id}", status_code=204)
async def delete(note_id: str, request: Request, _: SessionCookieDep, db: DatabaseDep):
    user = get_user_session_state(request)

    if user is None:
        raise HTTPException(status_code=401)

    delete_user_note(db, user.id, note_id)
    return PlainTextResponse(status_code=204)
