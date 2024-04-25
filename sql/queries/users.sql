-- name: CreateUser :one
INSERT INTO users (id, create_at, update_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
