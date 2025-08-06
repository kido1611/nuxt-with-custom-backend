using Backend.Application.DTOs;
using Backend.WebApi.Middlewares;

namespace Backend.WebApi.Routes;

public static class UserRoutes
{
    public static void MapUserRoutes(this WebApplication app)
    {
        app.MapGet("/api/user", Get)
            .WithName("Get User")
            .WithMetadata(new AuthMiddlewareAttribute());
    }

    private static IResult Get(HttpContext context)
    {
        var userResponse = context.Items["data_user_session"] as UserResponse;
        if (userResponse == null)
        {
            return Results.Json(new
            {
                message = "Unauthorized"
            }, statusCode: 401);
        }

        return Results.Json(new
        {
            data = userResponse
        });
    }
}