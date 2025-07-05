namespace Backend.Infrastructure.Configurations;

public class DatabaseSetting
{
    public string ConnectionString { get; init; } = "";
    public bool AutoMigration { get; init; }
}