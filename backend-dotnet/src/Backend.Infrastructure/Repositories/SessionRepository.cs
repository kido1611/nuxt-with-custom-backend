using Backend.Domain.Entities;
using Backend.Domain.Repositories;
using Backend.Infrastructure.Data;
using Microsoft.EntityFrameworkCore;

namespace Backend.Infrastructure.Repositories;

public class SessionRepository(AppDbContext context): ISessionRepository
{
    public async Task<Session> CreateAsync(Session session)
    {
        context.Sessions.Add(session);
        await context.SaveChangesAsync();

        return session;
    }

    public async Task<Session?> GetByIdAsync(string id)
    {
        return await context.Sessions.FirstOrDefaultAsync(p => p.Id == id);
    }

    public async Task<Session> UpdateAsync(Session session)
    {
        context.Sessions.Update(session);
        await context.SaveChangesAsync();

        return session;
    }

    public async Task DeleteByIdAsync(string id)
    {
        var session = await GetByIdAsync(id);
        if (session != null)
        {
            context.Sessions.Remove(session);
            await context.SaveChangesAsync();
        }
    }
}