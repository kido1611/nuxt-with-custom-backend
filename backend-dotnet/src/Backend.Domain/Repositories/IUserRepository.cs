using Backend.Domain.Entities;

namespace Backend.Domain.Repositories;

public interface IUserRepository
{
    Task<User?> GetByEmailAsync(string email);
    Task<User?> GetByIdAsync(string id);
    Task<User> CreateAsync(User user);
}
