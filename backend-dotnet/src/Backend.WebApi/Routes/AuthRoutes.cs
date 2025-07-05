using Backend.Application.DTOs;
using Backend.Application.Requests;
using Backend.Application.UseCases.Sessions;
using Backend.Application.UseCases.Users;
using Backend.WebApi.Middlewares;

namespace Backend.WebApi.Routes;

public static class AuthRoutes
{
    public static void MapAuthRoutes(this WebApplication app)
    {
        app.MapPost("/api/auth/register", Register).WithName("Auth Register")
            .WithMetadata(new GuestMiddlewareAttribute());
        app.MapPost("/api/auth/login", Login).WithName("Auth Login")
            .WithMetadata(new GuestMiddlewareAttribute());
        app.MapDelete("/api/auth/logout", Logout).WithName("Auth Logout")
            .WithMetadata(new AuthMiddlewareAttribute());
    }

    private static async Task<IResult> Login(HttpContext context, LoginUseCase loginUseCase, DeleteSessionUseCase deleteSessionUseCase, LoginRequest request)
    {
        var loginResponse = await loginUseCase.Execute(request);
        
        if (loginResponse.Error != null)
        {
            return Results.Json(new
            {
                message = loginResponse.Error.Message
            }, statusCode: loginResponse.Error.StatusCode);
        }

        if (context.Items["data_session"] is SessionResponse existingSessionResponse)
        {
            // logger.LogDebug("[APP] request have session. delete it first.");
            await deleteSessionUseCase.Execute(existingSessionResponse.Id);
        }
        
        context.Items["data_session"] = loginResponse.Data.SessionResponse;
        
        return Results.Json(new
        {
            data = loginResponse.Data.UserResponse
        }, statusCode: 200);
    }
    
    private static async Task<IResult> Register(HttpContext context, CreateUserUseCase createUserUseCase, DeleteSessionUseCase deleteSessionUseCase, RegisterRequest request)
    {
        var userResponse = await createUserUseCase.Execute(request);
        if (userResponse.Error != null)
        {
            return Results.Json(new
            {
                message = userResponse.Error.Message
            }, statusCode: userResponse.Error.StatusCode);
        }

        if (context.Items["data_session"] is not SessionResponse existingSessionResponse) 
            return Results.Json(
                userResponse,
                statusCode: 201);
        
        await deleteSessionUseCase.Execute(existingSessionResponse.Id);
        context.Items["data_session"] = null;

        return Results.Json(new
        {
            data = userResponse.Data
        }, statusCode: 201);
    }

    private static async Task<IResult> Logout(HttpContext context, DeleteSessionUseCase deleteSessionUseCase)
    {
        if (context.Items["data_session"] is SessionResponse sessionResponse)
        {
            await deleteSessionUseCase.Execute(sessionResponse.Id);

            context.Items["data_session"] = null;
            context.Items["data_user_session"] = null;
        }

        return Results.NoContent();
    }
}