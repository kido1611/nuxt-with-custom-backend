using System.Security.Cryptography;

namespace Backend.Application.Shared;

public static class TokenGenerator
{
    public static string GenerateToken()
    {
        var rng = RandomNumberGenerator.Create();
    
        var randomBytes = new byte[32];
        rng.GetBytes(randomBytes);
        return Convert.ToBase64String(randomBytes)
            .Replace("+", "-")  // Make URL-safe
            .Replace("/", "_")  // Make URL-safe
            .Replace("=", "");  // Remove padding       
    }
}