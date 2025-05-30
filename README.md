# User Notes Applications

Write your note now!

This project is created for learning Nuxt as frontend and Laravel as Backend (BFF).

## Requirement

- This app is created to save a notes.
- Note only can be created when an user is exist or was logged in.
- User can create many notes, but note required an user.
- User can comment to the note.
- User can register with email and password or SSO with Google/Github.

### Note

- Note have **title** and **description**, and description format is markdown.
- Note is possible to have attachments.
- Note is accessed through URL with slug identifier
- Note have a possibility to be publicly visible, but private as default.
- Attachment should be stored through Object Storage/S3.
- Deleted note will be moved into `Trash`. When note is deleted inside `Trash`,
  it will be permanently deleted.
