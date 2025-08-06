namespace Backend.Application.Requests;

public class RegisterRequest
{
    public required string name { get; set; }
    public required string email { get; set; }
    public required string password { get; set; }
}