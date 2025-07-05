using Backend.Application.Interfaces;
using Backend.Domain.Entities;
using Backend.Domain.Repositories;
using Backend.Infrastructure.Repositories;

namespace Backend.Infrastructure.Services;

public class SessionService(ISessionRepository sessionRepository) : ISessionService
{
    public async Task<Session?> GetByIdAsync(string id)
    {
        return await sessionRepository.GetByIdAsync(id);
    }

    public async Task<Session> CreateAsync(Session session)
    {
        return await sessionRepository.CreateAsync(session);
    }

    public async Task<Session> UpdateAsync(Session session)
    {
        return await sessionRepository.UpdateAsync(session);
    }

    public async Task DeleteAsync(string id)
    {
        await sessionRepository.DeleteByIdAsync(id);
    }
}