namespace Backend.WebApi.Configurations;

public class CorsSetting
{
    public required string[] Origins { get; init; }
    public required string[] Methods { get; init; }
    public required string[] Headers { get; init; }
    public bool AllowCredentials { get; init; }
}