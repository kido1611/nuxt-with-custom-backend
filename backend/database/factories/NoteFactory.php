<?php

namespace Database\Factories;

use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Note>
 */
class NoteFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'user_id' => User::factory(),
            'title' => fake()->realText(),
            'description' => fake()->realText(4000),
            'visible_at' => null,
        ];
    }

    public function public(): NoteFactory|Factory
    {
        return $this->state(fn () => [
            'visible_at' => now(),
        ]);
    }
}
