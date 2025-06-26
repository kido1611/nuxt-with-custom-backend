from fastapi import APIRouter, Depends, Request

from app.dependencies import SessionCookieDep, get_auth_route, get_user_session_state
from app.schemas.response_schema import ApiResponse
from app.schemas.user_schema import UserResponse

router = APIRouter(
    prefix="/user", tags=["user"], dependencies=[Depends(get_auth_route)]
)


@router.get("", status_code=200, response_model=ApiResponse[UserResponse])
@router.get("/", include_in_schema=False)
async def user(request: Request, _: SessionCookieDep):
    user = get_user_session_state(request)
    return ApiResponse(data=user)
