using Backend.Application.Requests;
using FluentValidation;

namespace Backend.Application.Validations;

public class CreateNoteRequestValidator: AbstractValidator<CreateNoteRequest>
{
    public CreateNoteRequestValidator()
    {
        RuleFor(x => x.Title).NotEmpty().WithMessage("Harus diisi.").MaximumLength(200);

        RuleFor(x => x.Description).MaximumLength(2000);
    }
}