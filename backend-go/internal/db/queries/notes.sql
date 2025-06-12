-- name: ListUserNotes :many
SELECT * FROM notes
WHERE user_id = ? ORDER BY created_at desc;

-- name: GetUserNote :one
SELECT * FROM notes
WHERE user_id = ? and id = ?;

-- name: CreateUserNote :one
INSERT INTO notes (id, user_id, title, description, visible_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateUserNote :one
UPDATE notes
SET title = ?, description = ?
WHERE id = ? and user_id = ?
RETURNING *;

-- name: DeleteUserNote :exec
DELETE FROM notes
WHERE id = ? and user_id = ?;
