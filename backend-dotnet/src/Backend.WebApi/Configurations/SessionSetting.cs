namespace Backend.WebApi.Configurations;

public class SessionSetting
{
    public string Domain { get; init; } = "localhost";
    public int Lifetime { get; init; }
}