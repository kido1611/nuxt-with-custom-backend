using Backend.Domain.Entities;
using Backend.Domain.Repositories;
using Backend.Infrastructure.Data;
using Microsoft.EntityFrameworkCore;

namespace Backend.Infrastructure.Repositories;

public class NoteRepository(AppDbContext context): INoteRepository
{
    public async Task<Note> CreateAsync(Note note)
    {
        context.Notes.Add(note);
        await context.SaveChangesAsync();

        return note;
    }

    public async Task<Note?> GetByIdAsync(string id)
    {
        var guid = Guid.Parse(id);
        return await context.Notes.FirstOrDefaultAsync(p => p.Id == guid);
    }

    public async Task<IEnumerable<Note>> GetAllAsync(string userId)
    {
        var userGuid = Guid.Parse(userId);
        return await context.Notes.Where(x => x.UserId == userGuid).ToListAsync();
    }

    public async Task DeleteByIdAsync(string userId, string id)
    {
        var noteGuid = Guid.Parse(id);
        var userGuid = Guid.Parse(userId);
        var note = await context.Notes.FirstOrDefaultAsync(p => p.UserId == userGuid && p.Id == noteGuid);
        if (note != null)
        {
            context.Notes.Remove(note);
            await context.SaveChangesAsync();
        }
    }
}