using Microsoft.EntityFrameworkCore;
using Backend.Infrastructure.Data;
using Microsoft.Extensions.DependencyInjection;

namespace Backend.Infrastructure;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(
        this IServiceCollection services,
        string connectionString
        )
    {
        services.AddDbContext<AppDbContext>(options => options.UseSqlite(connectionString));

        return services;
    }
}
