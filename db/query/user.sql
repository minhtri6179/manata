-- name: CreateUser :one
INSERT INTO "user" (
        user_name,
        hashed_password,
        first_name,
        last_name,
        date_of_birth,
        email
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM "user"
WHERE user_name = $1;
-- name: UpdateUser :exec
UPDATE "user"
SET email = $1,
    first_name = $2,
    last_name = $3
WHERE user_name = $4;
-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_name = $1;