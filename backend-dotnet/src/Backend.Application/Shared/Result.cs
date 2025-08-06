namespace Backend.Application.Shared;

public class ResultError(string message, int statusCode)
{
    public int StatusCode { get; } = statusCode;
    public string Message { get; } = message;

    public static ResultError Unauthorized() => new ("Unauthorized", 401);
    public static ResultError Forbidden() => new ("Forbidden", 403);
    public static ResultError NotFound() => new("Not Found", 404);
}

public class Result<T> (T data, ResultError? error)
{
    public T Data { get; set; } = data;
    public ResultError? Error { get; set; } = error;

    public static Result<T> Success(T data) => new(data, null);
    public static Result<T> Unauthorized() => new(default!, ResultError.Unauthorized());
    public static Result<T> Forbidden() => new(default!, ResultError.Forbidden());
    public static Result<T> NotFound() => new(default!, ResultError.NotFound());
}