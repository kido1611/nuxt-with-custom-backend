using Backend.Application;
using Microsoft.EntityFrameworkCore;
using Backend.Infrastructure;
using Backend.Infrastructure.Data;
using Backend.WebApi.Configurations;
using Backend.WebApi.Middlewares;
using Backend.WebApi.Routes;
using Serilog;
using Serilog.Core;
using Serilog.Events;
using static System.Boolean;

var logLevelSwitch = new LoggingLevelSwitch
{
    MinimumLevel = LogEventLevel.Debug
};

var logLevelFromEnv = Environment.GetEnvironmentVariable("APP_LOG_LEVEL");
if (Enum.TryParse<LogEventLevel>(logLevelFromEnv, true, out var parsedLevel))
{
    logLevelSwitch.MinimumLevel = parsedLevel;
}

Log.Logger = new LoggerConfiguration()
    .WriteTo.Console()
    .MinimumLevel.ControlledBy(logLevelSwitch)
    .CreateLogger();

var builder = WebApplication.CreateBuilder(args);

// Configurations
builder.Services.Configure<SessionSetting>(builder.Configuration.GetSection("Session"));
builder.Services.Configure<CorsSetting>(builder.Configuration.GetSection("Cors"));

builder.Host.UseSerilog();
builder.Services.AddInfrastructure(builder.Configuration);
builder.Services.AddApplication();

// Add services to the container.
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();
builder.Services.AddHealthChecks();

builder.Services.AddCors(options =>
{
    var corsSettings = builder.Configuration.GetSection("Cors").Get<CorsSetting>();
    if (corsSettings == null) return;

    options.AddPolicy(name: "CustomCors", policy =>
    {
        policy.WithOrigins(corsSettings.Origins)
            .WithHeaders(corsSettings.Headers)
            .WithMethods(corsSettings.Methods);
        if (corsSettings.AllowCredentials)
        {
            policy.AllowCredentials();
        }
    });
});

var app = builder.Build();

TryParse(builder.Configuration.GetSection("Database")["AutoMigrations"], out var result);
if (result)
{
    using var scope = app.Services.CreateScope();
    var db = scope.ServiceProvider.GetRequiredService<AppDbContext>();
    db.Database.Migrate();
}

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

app.UseCors("CustomCors");

app.UseMiddleware<OriginMiddleware>();
app.UseMiddleware<SessionMiddleware>();
app.UseMiddleware<CsrfMiddleware>();
app.UseMiddleware<GuestMiddleware>();
app.UseMiddleware<AuthMiddleware>();

app.MapSanctumRoutes();
app.MapHealthRoutes();
app.MapAuthRoutes();
app.MapUserRoutes();
app.MapNoteRoutes();

app.MapHealthChecks("/healthz");

app.Run();