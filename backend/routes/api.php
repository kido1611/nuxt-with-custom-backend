<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::apiResource('notes', \App\Http\Controllers\Api\NoteController::class)
    ->middleware(['auth:sanctum']);

Route::post('/auth/register', \App\Http\Controllers\Api\RegisterController::class)
    ->name('auth.register')
    ->middleware(['guest']);
Route::post('/auth/login', \App\Http\Controllers\Api\LoginController::class)
    ->name('auth.login')
    ->middleware(['guest']);
