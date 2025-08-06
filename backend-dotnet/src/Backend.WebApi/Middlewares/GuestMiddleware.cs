using Backend.Application.DTOs;

namespace Backend.WebApi.Middlewares;

public class GuestMiddleware(RequestDelegate next)
{
    public async Task InvokeAsync(HttpContext context)
    {
        var endpoint = context.GetEndpoint();
        if (endpoint?.Metadata.GetMetadata<GuestMiddlewareAttribute>() != null)
        {
            if (context.Items["data_user_session"] is UserResponse)
            {
                context.Response.StatusCode = 403;
                await context.Response.WriteAsync("Forbidden");
                
                return;
            }
        }
        
        await next(context);
    }
}

public class GuestMiddlewareAttribute : Attribute {}