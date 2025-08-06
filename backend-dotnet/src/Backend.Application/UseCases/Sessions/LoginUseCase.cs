using System.Security.Cryptography;
using Backend.Application.DTOs;
using Backend.Application.Interfaces;
using Backend.Application.Mapper;
using Backend.Application.Requests;
using Backend.Application.Shared;
using Backend.Domain.Entities;
using Backend.Domain.Repositories;

namespace Backend.Application.UseCases.Sessions;

public class LoginUseCase(IUserRepository userRepository, ISessionService sessionService, IPasswordService passwordService)
{
    public async Task<Result<LoginResponse>> Execute(LoginRequest request)
    {
        // TODO: Validate
        
        var user = await userRepository.GetByEmailAsync(request.email);
        if (user == null)
        {
            return Result<LoginResponse>.Unauthorized();
        }

        var isMatch = passwordService.VerifyPassword(request.password, user.Password);
        if (!isMatch)
        {
            return Result<LoginResponse>.Unauthorized();
        }
        
        var session = await sessionService.CreateAsync(new Session
        {
            Id = TokenGenerator.GenerateToken(),
            UserId = user.Id,
            CsrfToken = TokenGenerator.GenerateToken(),
            IpAddress = null,
            UserAgent = null,
            ExpiredAt = DateTime.Now.AddMinutes(180), // TODO: change 90 with value from settings
        });

        var sessionDto = SessionMapper.ToDto(session);
        var userDto = UserMapper.ToDto(user);

        var loginResponse = new LoginResponse
        {
            SessionResponse = sessionDto,
            UserResponse = userDto
        };
        
        return Result<LoginResponse>.Success(loginResponse);
    }
}