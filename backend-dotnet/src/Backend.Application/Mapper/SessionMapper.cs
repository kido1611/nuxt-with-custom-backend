using Backend.Application.DTOs;
using Backend.Domain.Entities;

namespace Backend.Application.Mapper;

public static class SessionMapper
{
    public static SessionResponse ToDto(Session session)
    {
        return new SessionResponse
        {
            Id = session.Id,
            UserId = session.UserId,
            CsrfToken =session.CsrfToken, 
            IpAddress = session.IpAddress,
            UserAgent = session.UserAgent,
            ExpiredAt = session.ExpiredAt,
        };
    }
}