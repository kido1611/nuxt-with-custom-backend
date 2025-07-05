namespace Backend.WebApi.Routes;

public static class HealthRoutes
{
    public static void MapHealthRoutes(this WebApplication app)
    {
        app.MapGet("/health", Get)
            .WithName("Health check");
    }

    private static IResult Get()
    {
        var response = new
        {
            message = "Alive"
        };
        return Results.Json(response);
    }
}