using Backend.Domain.Entities;

namespace Backend.Domain.Repositories;

public interface ISessionRepository
{
    Task<Session> CreateAsync(Session session);
    Task<Session?> GetByIdAsync(string id);
    Task<Session> UpdateAsync(Session session);
    Task DeleteByIdAsync(string id);
}
