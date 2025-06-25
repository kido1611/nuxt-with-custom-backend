from datetime import datetime, timedelta, timezone
from sqlmodel import Session
from starlette.middleware.base import BaseHTTPMiddleware, RequestResponseEndpoint
from starlette.requests import Request
from starlette.responses import Response, PlainTextResponse

from app.core.config import Config, get_config
from app.core.database import engine
from app.dependencies import get_session_state
from app.usecase.session_usecase import extend_expired_session, validate_session


class OriginMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next) -> Response:
        origin = request.headers.get("Origin")
        if origin is None:
            origin = request.headers.get("Referer")

        config = get_config()
        if not config.validate_origin(origin):
            return PlainTextResponse(status_code=403)

        return await call_next(request)


class SessionMiddleware(BaseHTTPMiddleware):
    async def dispatch(
        self, request: Request, call_next: RequestResponseEndpoint
    ) -> Response:
        session_id_cookie = request.cookies.get("app_session")
        if session_id_cookie is not None:
            with Session(engine) as session:
                data = validate_session(session, session_id_cookie)
                if data[0] is not None:
                    request.state.session = data[0]
                    request.state.session_user = data[1]

        response = await call_next(request)

        config = get_config()

        current_session = get_session_state(request)

        # clear cookie when user have session_cookie, but missing after request
        if current_session is None:
            if session_id_cookie is not None:
                clear_cookie_response(
                    response=response, config=config, key="app_session"
                )

            return response

        # set new cookie when old session id != new session id (after request /sanctum/csrf-cookie or after login)
        if session_id_cookie != current_session.id:
            response.set_cookie(
                key="app_session",
                value=current_session.id,
                expires=current_session.expired_at.astimezone(timezone.utc),
                path="/",
                domain=config.session_domain,
                secure=config.session_secure,
                httponly=True,
                samesite="lax",
            )
            return response

        if current_session.expired_at < (datetime.now() + timedelta(minutes=30)):
            with Session(engine) as session:
                updated_session = extend_expired_session(
                    session, config, current_session.id
                )
                if updated_session is not None:
                    response.set_cookie(
                        key="app_session",
                        value=current_session.id,
                        expires=current_session.expired_at.astimezone(timezone.utc),
                        path="/",
                        domain=config.session_domain,
                        secure=config.session_secure,
                        httponly=True,
                        samesite="lax",
                    )

        return response


class CsrfMiddleware(BaseHTTPMiddleware):
    async def dispatch(
        self, request: Request, call_next: RequestResponseEndpoint
    ) -> Response:
        not_safe_methods = ["post", "delete", "put", "patch"]
        session = get_session_state(request)

        method = request.method.lower()

        if method in not_safe_methods:
            csrfToken = request.headers.get("X-XSRF-TOKEN")
            if csrfToken is None:
                return PlainTextResponse(status_code=419)

            if session is None:
                return PlainTextResponse(status_code=419)

            if session.csrf_token != csrfToken:
                return PlainTextResponse(status_code=419)

        response = await call_next(request)

        new_session = get_session_state(request)
        config = get_config()

        # when logout or clear session after register
        if new_session is None:
            if session is not None:
                clear_cookie_response(
                    response=response, config=config, key="XSRF-TOKEN"
                )
        else:
            response.set_cookie(
                key="XSRF-TOKEN",
                value=new_session.csrf_token,
                expires=new_session.expired_at.astimezone(timezone.utc),
                path="/",
                domain=config.session_domain,
                secure=config.session_secure,
                httponly=False,
                samesite="lax",
            )

        return response


class CorsMiddleware(BaseHTTPMiddleware):
    config = get_config()

    async def dispatch(
        self, request: Request, call_next: RequestResponseEndpoint
    ) -> Response:
        if self.is_preflight_request(request):
            headers = [("Vary", "Origin")]
            if not self.is_origin_header(request, None):
                return PlainTextResponse(status_code=403, headers=dict(headers))

            origin = str(request.headers.get("Origin"))
            headers.append(("Access-Control-Allow-Origin", origin))

            if self.config.cors_allow_credentials:
                headers.append(("Access-Control-Allow-Credentials", "true"))

            headers.append(("Access-Control-Allow-Methods", self.config.cors_methods))
            headers.append(("Access-Control-Allow-Headers", self.config.cors_headers))

            return PlainTextResponse(status_code=204, headers=dict(headers))

        response = await call_next(request)

        if self.is_origin_header(request, response):
            if self.config.cors_allow_credentials:
                response.headers.append("Access-Control-Allow-Credentials", "true")

        return response

    def is_preflight_request(self, request: Request) -> bool:
        method = request.method.lower()
        accessControllRequestMethod = request.headers.get(
            "Access-Control-Request-Method"
        )

        return method == "options" and accessControllRequestMethod is not None

    def is_origin_header(self, request: Request, response: Response | None) -> bool:
        if response is not None:
            response.headers.add_vary_header("Origin")

        origin = request.headers.get("Origin")
        if origin is None:
            return False

        if not self.config.validate_origin(origin):
            return False

        if response is not None:
            response.headers.append("Access-Control-Allow-Origin", origin)

        return True


def clear_cookie_response(*, response: Response, config: Config, key: str):
    expires = datetime.now() - timedelta(days=1)
    response.set_cookie(
        key=key,
        value="",
        expires=expires.astimezone(timezone.utc),
        path="/",
        domain=config.session_domain,
        secure=config.session_secure,
        httponly=True,
        samesite="lax",
    )
