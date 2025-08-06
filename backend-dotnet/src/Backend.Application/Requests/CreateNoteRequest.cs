namespace Backend.Application.Requests;

public class CreateNoteRequest
{
    public required string Title { get; init; }   
    public string? Description { get; set;  }
}