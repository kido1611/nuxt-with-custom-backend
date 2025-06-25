from fastapi import APIRouter, Depends, Request
from fastapi.responses import PlainTextResponse

from app.schemas.response_schema import ApiResponse
from app.schemas.user_schema import LoginRequest, RegisterRequest, UserResponse
from app.dependencies import (
    ConfigDep,
    DatabaseDep,
    SessionCookieDep,
    get_auth_route,
    get_guest_route,
    get_session_state,
)
from app.usecase.session_usecase import create_session, delete_session
from app.usecase.user_usecase import check_user, create_user

router = APIRouter(prefix="/auth", tags=["auth"])


@router.post(
    "/login",
    status_code=200,
    response_model=ApiResponse[UserResponse],
    dependencies=[Depends(get_guest_route)],
)
async def login(
    login_request: LoginRequest, request: Request, db: DatabaseDep, config: ConfigDep
):
    user_response = check_user(db, login_request)

    current_session = get_session_state(request)
    if current_session is not None:
        delete_session(db, current_session.id)

    session_response = create_session(db, config, user_response)
    request.state.session = session_response
    request.state.session_user = user_response

    return ApiResponse(data=user_response)


@router.post(
    "/register",
    status_code=201,
    response_model=ApiResponse[UserResponse],
    dependencies=[Depends(get_guest_route)],
)
async def register(
    register_request: RegisterRequest,
    request: Request,
    db: DatabaseDep,
):
    user = create_user(db, register_request)

    current_session = get_session_state(request)
    if current_session is not None:
        delete_session(db, current_session.id)

    request.state.session = None
    request.state.session_user = None

    return ApiResponse(data=user)


@router.delete(
    "/logout",
    status_code=204,
    response_model=None,
    dependencies=[Depends(get_auth_route)],
)
async def logout(request: Request, db: DatabaseDep, _: SessionCookieDep):
    session_response = get_session_state(request)
    if session_response is not None:
        delete_session(db=db, session_id=session_response.id)

    request.state.session = None
    request.state.session_user = None

    return PlainTextResponse(status_code=204)
