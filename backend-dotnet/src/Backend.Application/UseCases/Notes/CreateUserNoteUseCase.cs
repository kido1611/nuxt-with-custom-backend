using Backend.Application.DTOs;
using Backend.Application.Mapper;
using Backend.Application.Requests;
using Backend.Application.Shared;
using Backend.Domain.Entities;
using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Notes;

public class CreateUserNoteUseCase(INoteRepository noteRepository)
{
    public async Task<Result<NoteResponse>> Execute(string userId, CreateNoteRequest request)
    {
        var note = new Note
        {
            Id = Guid.CreateVersion7(),
            UserId = Guid.Parse(userId),
            Title = request.Title,
            Description = request.Description,
        };

        var newNote = await noteRepository.CreateAsync(note);
        var noteResponse = NoteMapper.ToDto(newNote);

        return Result<NoteResponse>.Success(noteResponse);
    }
}