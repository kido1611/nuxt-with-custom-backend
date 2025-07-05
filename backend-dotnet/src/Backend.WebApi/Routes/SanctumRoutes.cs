using Backend.Application.UseCases.Sessions;

namespace Backend.WebApi.Routes;

public static class SanctumRoutes
{
    public static void MapSanctumRoutes(this WebApplication app)
    {
        app.MapGet("/sanctum/csrf-cookie", Get).WithName("Create temporary session");
    }

    private static async Task<IResult> Get(HttpContext context, CreateEmptySessionUseCase createEmptySessionUseCase)
    {
        if (context.Items["data_session"] != null) return Results.NoContent();
    
        var sessionResponse = await createEmptySessionUseCase.Execute();
        if (sessionResponse.Error == null)
        {
            context.Items["data_session"] = sessionResponse.Data;
        }

        return Results.NoContent();
    }
    
}