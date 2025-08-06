using Backend.Application.DTOs;

namespace Backend.WebApi.Middlewares;

public class AuthMiddleware(RequestDelegate next)
{
    public async Task InvokeAsync(HttpContext context)
    {
        var endpoint = context.GetEndpoint();
        if (endpoint?.Metadata.GetMetadata<AuthMiddlewareAttribute>() != null)
        {
            if (context.Items["data_user_session"] is not UserResponse)
            {
                context.Response.StatusCode = 401;
                await context.Response.WriteAsync("Unauthorized");
        
                return;
            }
        }

        await next(context);
    }   
}

public class AuthMiddlewareAttribute : Attribute
{
}