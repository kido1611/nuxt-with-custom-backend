<?php

namespace App\Http\Requests;

use App\Models\Note;
use Illuminate\Foundation\Http\FormRequest;

class NoteCreateRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return $this->user()->can('create', Note::class);
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'title' => ['required', 'string', 'ascii', 'max:200'],
            'description' => ['nullable', 'string', 'max:20000'],
            'visible_at' => ['nullable', 'date'],
        ];
    }
}
