using Backend.Application.DTOs;
using Backend.Application.Interfaces;
using Backend.Application.Mapper;
using Backend.Application.Shared;

namespace Backend.Application.UseCases.Sessions;

public class ExtendSessionUseCase(ISessionService sessionService)
{
    public async Task<Result<SessionResponse>> Execute(string sessionId)
    {
        var session = await sessionService.GetByIdAsync(sessionId);
        if (session == null)
        {
            return Result<SessionResponse>.Unauthorized();
        }

        if (session.ExpiredAt > DateTime.Now.AddMinutes(-30))
        {
            return Result<SessionResponse>.Forbidden();
        }
        
        session.ExpiredAt = DateTime.Now.AddMinutes(180);
        await sessionService.UpdateAsync(session);

        var sessionResponse = SessionMapper.ToDto(session);
        
        return Result<SessionResponse>.Success(sessionResponse);
    }
}