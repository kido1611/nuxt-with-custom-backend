using Backend.WebApi.Configurations;
using Microsoft.Extensions.Options;

namespace Backend.WebApi.Middlewares;

public class OriginMiddleware(RequestDelegate next, ILogger<OriginMiddleware> logger, IOptions<CorsSetting> corsSetting)
{
    public async Task InvokeAsync(HttpContext context)
    {
        var origin = context.Request.Headers["Origin"].ToString();
        
        if (string.IsNullOrEmpty(origin))
        {
            origin = context.Request.Headers["Referer"].ToString();
        }
        
        logger.LogDebug("[APP] Request origin: " + origin);

        if (string.IsNullOrEmpty(origin))
        {
            context.Response.StatusCode = 403;
            await context.Response.WriteAsync("Forbidden");
        
            return;
        }

        if (!IsOriginExist(origin))
        {
            logger.LogDebug("[APP] origin not listed");
            
            context.Response.StatusCode = 403;
            await context.Response.WriteAsync("Forbidden");
        
            return;
        }
        
        await next(context);
    }

    private bool IsOriginExist(string origin)
    {
        if (string.IsNullOrEmpty(origin))
        {
            return false;
        }
        
        var uri = new Uri(origin);
        var realOrigin = uri.Scheme + "://" + uri.Authority;

        return corsSetting.Value.Origins.Any(x =>
        {
            var uriCors = new Uri(x);
            var originCors = uriCors.Scheme + "://" + uriCors.Authority;
            
            logger.LogDebug("[APP] Comparing: " + realOrigin + "==" + originCors);
            
            return realOrigin == originCors;
        });
    }
}