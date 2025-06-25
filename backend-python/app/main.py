from fastapi import FastAPI
from app.routers import health, auth, note, sanctum, user

app = FastAPI()
app.include_router(health.router)
app.include_router(sanctum.router)
app.include_router(auth.router, prefix="/api")
app.include_router(note.router, prefix="/api")
app.include_router(user.router, prefix="/api")
