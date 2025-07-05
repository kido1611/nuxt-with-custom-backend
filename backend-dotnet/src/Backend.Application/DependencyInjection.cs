using Backend.Application.UseCases.Notes;
using Backend.Application.UseCases.Sessions;
using Backend.Application.UseCases.Users;
using FluentValidation.AspNetCore;
using Microsoft.Extensions.DependencyInjection;

namespace Backend.Application;

public static class DependencyInjection
{
    public static IServiceCollection AddApplication(this IServiceCollection services)
    {
        services.AddScoped<LoginUseCase>();
        services.AddScoped<CheckSessionUseCase>();
        services.AddScoped<CreateEmptySessionUseCase>();
        services.AddScoped<ExtendSessionUseCase>();
        services.AddScoped<DeleteSessionUseCase>();
        
        services.AddScoped<GetUserUseCase>();
        services.AddScoped<CreateUserUseCase>();
        
        services.AddScoped<ListUserNotesUseCase>();
        services.AddScoped<CreateUserNoteUseCase>();
        services.AddScoped<DeleteUserNoteUseCase>();

        services.AddFluentValidationAutoValidation();
        
        return services;
    }
}