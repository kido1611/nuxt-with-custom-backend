from fastapi import FastAPI, Request
from fastapi.exceptions import RequestValidationError
from fastapi.responses import PlainTextResponse
from app.core import middleware
from app.routers import health, auth, note, sanctum, user
import app.models

app = FastAPI()


@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, exc: RequestValidationError):
    return PlainTextResponse(status_code=400)


app.include_router(health.router)
app.include_router(sanctum.router)
app.include_router(auth.router, prefix="/api")
app.include_router(note.router, prefix="/api")
app.include_router(user.router, prefix="/api")

app.add_middleware(middleware.CsrfMiddleware)
app.add_middleware(middleware.SessionMiddleware)
app.add_middleware(middleware.OriginMiddleware)
app.add_middleware(middleware.CorsMiddleware)
