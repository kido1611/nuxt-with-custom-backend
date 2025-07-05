using Backend.Application.DTOs;
using Backend.Application.Interfaces;
using Backend.Application.Mapper;
using Backend.Application.Requests;
using Backend.Application.Shared;
using Backend.Domain.Entities;
using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Users;

public class CreateUserUseCase(IUserRepository userRepository, IPasswordService passwordService)
{
    public async Task<Result<UserResponse>> Execute(RegisterRequest request)
    {
        var userByEmail = await userRepository.GetByEmailAsync(request.email);
        if (userByEmail != null)
        {
            return new Result<UserResponse>(
                null!,
                new ResultError("Email was used.", 422)
            );
        }

        var user = new User
        {
            Id = Guid.CreateVersion7(),
            Email = request.email,
            Name = request.name,
            Password = passwordService.HashPassword(request.password)
        };

        var newUser = await userRepository.CreateAsync(user);
        var userResponse = UserMapper.ToDto(newUser);
        return Result<UserResponse>.Success(userResponse);
    }
}