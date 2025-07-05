using Backend.Domain.Entities;

namespace Backend.Application.Interfaces;

public interface ISessionService
{
    Task<Session?> GetByIdAsync(string id);
    Task<Session> CreateAsync(Session session);
    Task<Session> UpdateAsync(Session session);
    Task DeleteAsync(string id);
}