namespace Backend.Application.DTOs;

public class UserResponse
{
    public required string Id { get; set; }
    public required string Name { get; set; }
    public required string Email { get; set; }
    public DateTime CreatedAt { get; set; }
}