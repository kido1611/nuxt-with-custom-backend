using Backend.Application.DTOs;
using Backend.WebApi.Configurations;
using Microsoft.Extensions.Options;

namespace Backend.WebApi.Middlewares;

public class CsrfMiddleware(RequestDelegate next, ILogger<CsrfMiddleware> logger, IOptions<SessionSetting> sessionSetting)
{
    public async Task InvokeAsync(
        HttpContext context )
    {
        logger.LogDebug("[APP] starting csrf middleware");
        
        var method = context.Request.Method;
        var unsafeMethods = new List<string> { "POST", "PUT", "PATCH", "DELETE" };
        var isUnsafeMethod = unsafeMethods.Contains(method);
        var sessionResponse = context.Items["data_session"] as SessionResponse;

        if (isUnsafeMethod)
        {
            var csrfToken = context.Request.Headers["X-XSRF-TOKEN"].ToString();
            if (string.IsNullOrEmpty(csrfToken) || sessionResponse == null || sessionResponse.CsrfToken != csrfToken)
            {
                context.Response.StatusCode = 419;
                await context.Response.WriteAsync("CSRF Token Miss Match");

                return;
            }
        }
        
        context.Response.OnStarting(() =>
        {
            if (context.Items["data_session"] is SessionResponse newSessionResponse)
            {
                logger.LogDebug("[APP] add csrf token cookie");
                context.Response.Cookies.Append("XSRF-TOKEN", newSessionResponse.CsrfToken, new CookieOptions
                {
                    HttpOnly = false,
                    Secure = true,
                    SameSite = SameSiteMode.Lax,
                    Path = "/",
                    Domain = sessionSetting.Value.Domain,
                    Expires = newSessionResponse.ExpiredAt,
                });
            }
            else if(sessionResponse != null) 
            {
                logger.LogDebug("[APP] delete csrf cookie because session is missing from context");
                
                context.Response.Cookies.Append("XSRF-TOKEN", "", new CookieOptions
                {
                    HttpOnly = false,
                    Secure = true,
                    SameSite = SameSiteMode.Lax,
                    Path = "/",
                    Domain = sessionSetting.Value.Domain,
                    Expires = DateTime.Now.AddDays(-1),
                });
            }
            
            logger.LogDebug("[APP] end csrf middleware");
            return Task.CompletedTask;
        });
        await next(context);
    }
}