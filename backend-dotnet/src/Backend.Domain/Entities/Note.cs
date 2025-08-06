namespace Backend.Domain.Entities;

public class Note
{
    public required Guid Id { get; set; }
    public required Guid UserId { get; set; }
    public required string Title { get; set; }
    public string? Description { get; set; } = string.Empty;
    public DateTime? VisibleAt { get; set; } = null;
    public DateTime CreatedAt { get; set; } = DateTime.Now;
    public DateTime UpdatedAt { get; set; } = DateTime.Now;
    public DateTime? DeletedAt { get; set; } = null;


    public User User { get; set; } = null!;
}
