using Backend.Domain.Entities;

namespace Backend.Domain.Repositories;

public interface INoteRepository
{
    Task<Note> CreateAsync(Note note);
    Task<Note?> GetByIdAsync(string id);
    Task<IEnumerable<Note>> GetAllAsync(string userId);
    Task DeleteByIdAsync(string userId, string id);
}
