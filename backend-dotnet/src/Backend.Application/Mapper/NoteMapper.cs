using Backend.Application.DTOs;
using Backend.Domain.Entities;

namespace Backend.Application.Mapper;

public static class NoteMapper
{
    public static NoteResponse ToDto(Note note)
    {
        return new NoteResponse
        {
            Id = note.Id.ToString(),
            Title = note.Title,
            Description = note.Description,
            UserId = note.UserId.ToString(),
            CreatedAt = note.CreatedAt,
            VisibleAt = note.VisibleAt
        };
    }
}