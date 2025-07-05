using Backend.Application.DTOs;
using Backend.Application.Mapper;
using Backend.Application.Shared;
using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Notes;

public class ListUserNotesUseCase(INoteRepository noteRepository)
{
    public async Task<Result<List<NoteResponse>>> Execute(string userId)
    {
        var notes = await noteRepository.GetAllAsync(userId);
        var noteResponses = notes.Select(NoteMapper.ToDto);
        
        return Result<List<NoteResponse>>.Success(noteResponses.ToList());
    }
}