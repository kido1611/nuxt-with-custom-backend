using Backend.Application.DTOs;
using Backend.Domain.Entities;

namespace Backend.Application.Mapper;

public static class UserMapper
{
    public static UserResponse ToDto(User user)
    {
        return new UserResponse
        {
            Id = user.Id.ToString(),
            Name = user.Name,
            Email = user.Email,
            CreatedAt = user.CreatedAt
        };
    }
}