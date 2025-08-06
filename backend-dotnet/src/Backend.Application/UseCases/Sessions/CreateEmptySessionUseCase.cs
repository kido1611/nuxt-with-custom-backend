using Backend.Application.DTOs;
using Backend.Application.Interfaces;
using Backend.Application.Mapper;
using Backend.Application.Shared;
using Backend.Domain.Entities;

namespace Backend.Application.UseCases.Sessions;

public class CreateEmptySessionUseCase(ISessionService sessionService)
{
    public async Task<Result<SessionResponse>> Execute()
    {
        var session = await sessionService.CreateAsync(new Session
        {
            Id = TokenGenerator.GenerateToken(),
            UserId = null,
            CsrfToken = TokenGenerator.GenerateToken(),
            IpAddress = null,
            UserAgent = null,
            ExpiredAt = DateTime.Now.AddMinutes(180), // TODO: change 90 with value from settings
        });

        var sessionDto = SessionMapper.ToDto(session);
        
        return Result<SessionResponse>.Success(sessionDto);
    }
}