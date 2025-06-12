-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users 
(id, name, email, password)
VALUES
(?, ?, ?, ?)
RETURNING *;
