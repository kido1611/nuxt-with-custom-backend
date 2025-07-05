using Backend.Application.Requests;
using FluentValidation;

namespace Backend.Application.Validations;

public class LoginRequestValidator: AbstractValidator<LoginRequest>
{
    public LoginRequestValidator()
    {
        RuleFor(x => x.email).NotEmpty().WithMessage("Harus diisi.")
            .EmailAddress();

        RuleFor(x => x.password).NotEmpty().WithMessage("Harus diisi.");
    }
}