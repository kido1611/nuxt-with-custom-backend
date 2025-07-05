using Backend.Application.DTOs;
using Backend.Application.Interfaces;
using Backend.Application.Mapper;
using Backend.Application.Shared;

namespace Backend.Application.UseCases.Sessions;

public class CheckSessionUseCase(ISessionService sessionService)
{
    public async Task<Result<SessionResponse>> Execute(string sessionId)
    {
        var session = await sessionService.GetByIdAsync(sessionId);
        if (session == null)
        {
            return Result<SessionResponse>.NotFound();
        }

        if (session.ExpiredAt < DateTime.Now)
        {
            return Result<SessionResponse>.Unauthorized();
        }

        var sessionResponse = SessionMapper.ToDto(session);
        
        return Result<SessionResponse>.Success(sessionResponse);
    }
}