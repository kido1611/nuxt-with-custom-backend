# User Notes Applications

Example of Nuxt (Frontend + Laravel (Backend), using Laravel Sanctum SPA mode. There are 2 example of frontend apps, 
first is using external module and second is manually by creating plugin (the reason is to understand how sanctum SPA mode work).

Another example can be found on this repository https://github.com/kido1611/laravel-breeze-nuxt-auth-example.

## Used library

- https://github.com/manchenkoff/nuxt-auth-sanctum/

## Step

- https://github.com/kido1611/laravel-nuxt-example/commit/785a13c103783bf268598c18d3ad6301e5d16ebd

### Frontend

- https://github.com/kido1611/laravel-nuxt-example/commit/476a04abe870b111d95650c644c166f3d198eee1
- https://github.com/kido1611/laravel-nuxt-example/commit/883f35720705407ef17f5a3d12110932dde2a642
- https://github.com/kido1611/laravel-nuxt-example/commit/a16d8aafc1c6c04f60db2b4b045410934611d73f
- https://github.com/kido1611/laravel-nuxt-example/commit/0134522e1057e5b23cf04cc18117604b13278e67

### Frontend 2

- https://github.com/kido1611/laravel-nuxt-example/commit/44e4eb4135bd91e613ddd0c6e10ca53d66fefc00

---

# Write your note now!

This project is created for learning Nuxt as frontend and Laravel as Backend (BFF).

## Requirement

- This app is created to save a notes.
- Note only can be created when an user is exist or was logged in.
- User can create many notes, but note required an user.
- User can comment to the note.
- User can register with email and password or SSO with Google/Github.

### Note

- Note have **title** and **description** ~~, and description format is markdown~~.
- ~~Note is possible to have attachments~~.
- Note is accessed through URL with slug identifier
- Note have a possibility to be publicly visible, but private as default.
- Attachment should be stored through Object Storage/S3.
- Deleted note will be moved into `Trash`. When note is deleted inside `Trash`,
  it will be permanently deleted.
