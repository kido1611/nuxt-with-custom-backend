using Backend.Application.Interfaces;
using Backend.Domain.Repositories;
using Backend.Infrastructure.Configurations;
using Microsoft.EntityFrameworkCore;
using Backend.Infrastructure.Data;
using Backend.Infrastructure.Repositories;
using Backend.Infrastructure.Services;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;

namespace Backend.Infrastructure;

public static class DependencyInjection
{
    public static IServiceCollection AddInfrastructure(
        this IServiceCollection services,
        IConfiguration configuration
    )
    {
        var settings = configuration
            .GetSection("Database")
            .Get<DatabaseSetting>();
        if (settings != null)
        {
            services.AddSingleton(settings);
        }
        services.AddDbContext<AppDbContext>((sp, options) =>
        {
            var setting = sp.GetRequiredService<DatabaseSetting>();
            options.UseSqlite(setting.ConnectionString);
        });

        services.AddScoped<IUserRepository, UserRepository>();
        services.AddScoped<ISessionRepository, SessionRepository>();
        services.AddScoped<INoteRepository, NoteRepository>();

        services.AddScoped<IPasswordService, PasswordService>();
        services.AddScoped<ISessionService, SessionService>();

        return services;
    }
}
