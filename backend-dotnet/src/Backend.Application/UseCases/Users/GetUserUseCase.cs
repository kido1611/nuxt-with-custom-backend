using System.Text.Json;
using Backend.Application.DTOs;
using Backend.Application.Mapper;
using Backend.Application.Shared;
using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Users;

public class GetUserUseCase(IUserRepository userRepository)
{
    public async Task<Result<UserResponse>> Execute(string userId)
    {
        var user = await userRepository.GetByIdAsync(userId);
        if (user == null)
        {
            return Result<UserResponse>.NotFound();
        }

        var userResponse = UserMapper.ToDto(user);
        return Result<UserResponse>.Success(userResponse);
    }
}