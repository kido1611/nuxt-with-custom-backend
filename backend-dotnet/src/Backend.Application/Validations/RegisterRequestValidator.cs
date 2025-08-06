using Backend.Application.Requests;
using FluentValidation;

namespace Backend.Application.Validations;

public class RegisterRequestValidator: AbstractValidator<RegisterRequest>
{
    public RegisterRequestValidator()
    {
        RuleFor(x => x.name).NotEmpty().WithMessage("Harus diisi.")
            .MaximumLength(100).WithMessage("Maksimal 100 huruf.");
        
        RuleFor(x => x.email).NotEmpty().WithMessage("Harus diisi.")
            .EmailAddress()
            .MaximumLength(100).WithMessage("Maksimal 100 huruf.");

        RuleFor(x => x.password).NotEmpty().WithMessage("Harus diisi.")
            .MaximumLength(100).WithMessage("Maksimal 100 huruf.");
    }
}