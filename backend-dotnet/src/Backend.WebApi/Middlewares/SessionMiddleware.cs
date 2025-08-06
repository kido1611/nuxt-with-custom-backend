using System.Text.Json;
using Backend.Application.DTOs;
using Backend.Application.UseCases.Sessions;
using Backend.Application.UseCases.Users;
using Backend.WebApi.Configurations;
using Microsoft.Extensions.Options;

namespace Backend.WebApi.Middlewares;

public class SessionMiddleware(RequestDelegate next, ILogger<SessionMiddleware> logger, IOptions<SessionSetting> sessionSetting)
{
    public async Task InvokeAsync(HttpContext context, CheckSessionUseCase checkSessionUseCase, ExtendSessionUseCase extendSessionUseCase, GetUserUseCase getUserUseCase)
    {
        var sessionId = context.Request.Cookies["app_session"];
        if (!string.IsNullOrEmpty(sessionId))
        {
            logger.LogDebug("[APP] Request have a session cookie");
            var result = await checkSessionUseCase.Execute(sessionId);
            if (result.Error != null)
            {
                logger.LogDebug("[APP] Clearing session cookie because session is missing/expired");
                
                context.Response.Cookies.Append("app_session", "", new CookieOptions
                {
                    HttpOnly = true,
                    Secure = true,
                    SameSite = SameSiteMode.Lax,
                    Path = "/",
                    Domain = sessionSetting.Value.Domain,
                    Expires = DateTime.Now.AddDays(-1),
                });
            }
            else
            {
                context.Items["data_session"] = result.Data;

                if (result.Data.UserId != null)
                {
                    logger.LogDebug(JsonSerializer.Serialize(result.Data));
                    logger.LogDebug(result.Data.UserId.ToString());
                    var user = await getUserUseCase.Execute(result.Data.UserId.ToString() ?? "");
                    if (user.Error == null)
                    {
                        logger.LogDebug("[APP] session have user.");
                        context.Items["data_user_session"] = user.Data;
                    }
                    else
                    {
                        logger.LogDebug("[APP] clear session because session bounded to user, but user is missing.");
                        context.Items["data_session"] = null;
                    }
                }
                else
                {
                    logger.LogDebug("[APP] session does not have user.");
                }
            }
        }
        else
        {
            logger.LogDebug("[APP] Request does not have a session cookie");
        }
        
        // must be used response on starting to allow set cookie after handling controller
        context.Response.OnStarting(async () =>
        {
            var sessionAfter = context.Items["data_session"] as SessionResponse;
            if (string.IsNullOrEmpty(sessionId) && sessionAfter != null)
            {
                logger.LogDebug("[APP] set cookie app_session because after logged in");
                logger.LogDebug(JsonSerializer.Serialize(sessionAfter));
                
                context.Response.Cookies.Append("app_session", sessionAfter.Id, new CookieOptions
                {
                    HttpOnly = true,
                    Secure = true,
                    SameSite = SameSiteMode.Lax,
                    Path = "/",
                    Domain = sessionSetting.Value.Domain,
                    Expires = sessionAfter.ExpiredAt,
                });
                return;
            }
            
            if (!string.IsNullOrEmpty(sessionId) && sessionAfter == null)
            {
                logger.LogDebug("[APP] delete cookie app_session because logged out");
                
                context.Response.Cookies.Append("app_session", "", new CookieOptions
                {
                    HttpOnly = true,
                    Secure = true,
                    SameSite = SameSiteMode.Lax,
                    Path = "/",
                    Domain = sessionSetting.Value.Domain,
                    Expires = DateTime.Now.AddDays(-1),
                });
                
                return;
            }
            
            if (!string.IsNullOrEmpty(sessionId) && sessionAfter != null)
            {
                if (sessionId != sessionAfter.Id)
                {
                    
                    logger.LogDebug("[APP] set cookie app_session because after logged in");
                    logger.LogDebug(JsonSerializer.Serialize(sessionAfter));
                
                    context.Response.Cookies.Append("app_session", sessionAfter.Id, new CookieOptions
                    {
                        HttpOnly = true,
                        Secure = true,
                        SameSite = SameSiteMode.Lax,
                        Path = "/",
                        Domain = sessionSetting.Value.Domain,
                        Expires = sessionAfter.ExpiredAt,
                    });
                }
                else
                {
                    
                    var extendSession = await extendSessionUseCase.Execute(sessionAfter.Id);
                    if (extendSession.Error == null)
                    {
                        logger.LogDebug("[APP] session still exist and almost expired");
                        context.Response.Cookies.Append("app_session", sessionAfter.Id, new CookieOptions
                        {
                            HttpOnly = true,
                            Secure = true,
                            SameSite = SameSiteMode.Lax,
                            Path = "/",
                            Domain = sessionSetting.Value.Domain,
                            Expires = sessionAfter.ExpiredAt,
                        });
                    }
                    else
                    {
                        logger.LogDebug("[APP] session still exist");
                    }
                }
                
                return ;
            }
            
            logger.LogDebug("[APP] do nothing with cookie (not logged in).");
        });

        await next(context);
    }
}
