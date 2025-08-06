using Backend.Application.DTOs;
using Backend.Application.Requests;
using Backend.Application.UseCases.Notes;
using Backend.WebApi.Middlewares;

namespace Backend.WebApi.Routes;

public static class NoteRoutes
{
    public static void MapNoteRoutes(this WebApplication app)
    {
        app.MapGet("/api/notes", List).WithName("List user notes")
            .WithMetadata(new AuthMiddlewareAttribute());
        app.MapPost("/api/notes", Create).WithName("Create user notes")
            .WithMetadata(new AuthMiddlewareAttribute());
        app.MapDelete("/api/notes/{noteId:guid}", Delete).WithName("Delete user note")
            .WithMetadata(new AuthMiddlewareAttribute());
    }
    private static async Task<IResult> List(HttpContext context, ListUserNotesUseCase listUserNotesUseCase)
    {
        var userResponse = context.Items["data_user_session"] as UserResponse;
        if (userResponse == null)
        {
            return Results.Json(new
            {
                message = "Unauthorized"
            }, statusCode: 401);
        }
        
        var notes = await listUserNotesUseCase.Execute(userResponse.Id);

        return Results.Json(new
        {
            data = notes.Data
        });
    }

    private static async Task<IResult> Create(HttpContext context, CreateNoteRequest request, CreateUserNoteUseCase createUserNoteUseCase)
    {
        var userResponse = context.Items["data_user_session"] as UserResponse;
        if (userResponse == null)
        {
            return Results.Json(new
            {
                message = "Unauthorized"
            }, statusCode: 401);
        }

        var noteResponse = await createUserNoteUseCase.Execute(userResponse.Id, request);
        if (noteResponse.Error != null)
        {
            return Results.Json(new
            {
                message = noteResponse.Error.Message
            }, statusCode: noteResponse.Error.StatusCode);
        }

        return Results.Json(new
        {
            data = noteResponse.Data
        }, statusCode: 201);
    }
    
    private static async Task<IResult> Delete(HttpContext context, Guid noteId,  DeleteUserNoteUseCase deleteUserNoteUseCase)
    {
        var userResponse = context.Items["data_user_session"] as UserResponse;
        if (userResponse == null)
        {
            return Results.Json(new
            {
                message = "Unauthorized"
            }, statusCode: 401);
        }

        await deleteUserNoteUseCase.Execute(userResponse.Id, noteId.ToString());

        return Results.NoContent();
    }
}