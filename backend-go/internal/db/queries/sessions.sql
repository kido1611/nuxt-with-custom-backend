-- name: CreateSession :one
INSERT INTO sessions (id, user_id, csrf_token, ip_address, user_agent, expired_at, last_activity_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetSessionById :one
SELECT * FROM sessions
WHERE id = ?
LIMIT 1;

-- name: UpdateSessionLastActivity :exec
UPDATE sessions
SET 
  last_activity_at = now()
WHERE id = ?;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;
