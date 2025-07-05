namespace Backend.Application.DTOs;

public class NoteResponse
{
    public required string Id { get; set; }
    public required string UserId { get; set; }
    public required string Title { get; set; }
    public string? Description { get; set; }
    public DateTime? VisibleAt { get; set; }
    public DateTime CreatedAt { get; set; }
}