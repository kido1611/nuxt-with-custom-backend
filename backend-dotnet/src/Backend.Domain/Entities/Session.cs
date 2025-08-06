namespace Backend.Domain.Entities;

public class Session
{
    public required string Id { get; set; }
    public Guid? UserId { get; set; }
    public required string CsrfToken { get; set; }
    public string? IpAddress { get; set; }
    public string? UserAgent { get; set; }
    public DateTime ExpiredAt { get; set; }
    public DateTime CreatedAt { get; set; } = DateTime.Now;
    public DateTime UpdatedAt { get; set; } = DateTime.Now;

    public User? User { get; set; } = null!;
}
