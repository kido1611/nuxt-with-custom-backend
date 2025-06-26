from fastapi import APIRouter, Request
from fastapi.responses import PlainTextResponse

from app.dependencies import ConfigDep, DatabaseDep, get_session_state
from app.usecase.session_usecase import create_session

router = APIRouter(tags=["auth"])


@router.get("/sanctum/csrf-cookie", status_code=204, response_model=None)
async def csrf_token(request: Request, db: DatabaseDep, config: ConfigDep):
    current_state = get_session_state(request)
    if current_state is not None:
        return PlainTextResponse(status_code=204)

    session_response = create_session(db, config, None)

    request.state.session = session_response

    return PlainTextResponse(status_code=204)
