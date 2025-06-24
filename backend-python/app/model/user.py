from pydantic import BaseModel, Field


class LoginRequest(BaseModel):
    email: str = Field(max_length=100)
    password: str = Field(max_length=100)


class RegisterRequest(LoginRequest):
    name: str = Field(max_length=100)
