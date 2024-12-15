-- name: CreateUser :one

INSERT INTO users (id, username, password, role) VALUES (
  gen_random_uuid(),
  $1,
  $2,
  $3
) RETURNING *;



