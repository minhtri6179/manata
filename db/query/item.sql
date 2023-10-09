-- name: GetTask :one
SELECT *
FROM task
WHERE id = $1;
-- name: CreateTask :one
INSERT INTO task (
        title,
        description,
        image,
        status
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: ListTasks :many
SELECT *
FROM task
ORDER BY id ASC
LIMIT $1 OFFSET $2;
-- name: UpdateStatus :exec
UPDATE task
SET title = $1,
    description = $2,
    status = $3
WHERE id = $4
RETURNING *;
-- name: DeleteTask :one
DELETE FROM task
WHERE id = $1
RETURNING *;