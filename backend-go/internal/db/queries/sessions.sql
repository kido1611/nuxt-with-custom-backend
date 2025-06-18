-- name: CreateSession :one
INSERT INTO sessions (id, user_id, csrf_token, ip_address, user_agent, expired_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetSessionById :one
SELECT * FROM sessions
WHERE id = ?
LIMIT 1;

-- name: UpdateSessionExpired :exec
UPDATE sessions
SET 
  expired_at = ?
WHERE id = ?;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;
