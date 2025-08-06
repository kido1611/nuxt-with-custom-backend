using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Notes;

public class DeleteUserNoteUseCase(INoteRepository noteRepository)
{
    public async Task Execute(string userId, string noteId)
    {
        await noteRepository.DeleteByIdAsync(userId, noteId);
    }
}