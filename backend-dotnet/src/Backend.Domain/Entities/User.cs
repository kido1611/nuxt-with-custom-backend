namespace Backend.Domain.Entities;

public class User
{
    public required Guid Id { get; set; }
    public required string Name { get; set; }
    public required string Email { get; set; }
    public required string Password { get; set; }
    public DateTime CreatedAt { get; set; } = DateTime.Now;
    public DateTime UpdatedAt { get; set; } = DateTime.Now;

    public ICollection<Note> Notes { get; } = new List<Note>();
    public ICollection<Session> Sessions { get; } = new List<Session>();
}
