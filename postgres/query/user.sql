-- name: GetUserByID :one
SELECT * FROM users
  WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
  WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    id, email
  ) VALUES (
    $1, $2
  )
  RETURNING *;
