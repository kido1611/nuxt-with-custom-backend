using Backend.Application.Interfaces;
using Backend.Application.Shared;

namespace Backend.Application.UseCases.Sessions;

public class DeleteSessionUseCase(ISessionService sessionService)
{
    public async Task Execute(string sessionId)
    {
        await sessionService.DeleteAsync(sessionId);
    }
}