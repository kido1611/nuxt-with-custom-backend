using Backend.Domain.Entities;
using Backend.Domain.Repositories;
using Backend.Infrastructure.Data;
using Microsoft.EntityFrameworkCore;

namespace Backend.Infrastructure.Repositories;

public class UserRepository(AppDbContext context) : IUserRepository
{
    public async Task<User?> GetByEmailAsync(string email)
    {
        return await context.Users
            .FirstOrDefaultAsync(p => p.Email == email);
    }

    public async Task<User?> GetByIdAsync(string id)
    {
        var guid = Guid.Parse(id);
        return await context.Users
            .FirstOrDefaultAsync(p => p.Id == guid);
    }

    public async Task<User> CreateAsync(User user)
    {
        context.Users.Add(user);
        await context.SaveChangesAsync();

        return user;
    }
}