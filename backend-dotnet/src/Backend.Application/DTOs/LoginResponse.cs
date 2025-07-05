namespace Backend.Application.DTOs;

public class LoginResponse
{
    public required SessionResponse SessionResponse { get; init; }
    public required UserResponse UserResponse { get; init; }
}