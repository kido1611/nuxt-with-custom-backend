namespace Backend.Application.DTOs;

public class SessionResponse
{
    public required string Id { get; init; }
    public Guid? UserId { get; init; }
    public required string CsrfToken { get; init; }
    public string? IpAddress { get; set; }
    public string? UserAgent { get; set; }
    public DateTime ExpiredAt { get; init; }
}