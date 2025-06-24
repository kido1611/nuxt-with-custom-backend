from pydantic import BaseModel, Field


class NoteRequest(BaseModel):
    title: str = Field(max_length=200)
    description: str | None = Field(max_length=2000)
