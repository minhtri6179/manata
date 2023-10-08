-- name CreateTask :one
INSERT INTO tasks (
        name,
        description,
        status,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING (
        id,
        name,
        description,
        status,
        created_at,
        updated_at
    );
-- name GetTask :one
SELECT id,
    name,
    description,
    status,
    created_at,
    updated_at
FROM tasks
WHERE id = $1;
-- name ListTasks :many
SELECT id,
    name,
    description,
    status,
    created_at,
    updated_at
FROM tasks
ORDER BY id ASC
LIMIT $1 OFFSET $2;
-- name UpdateTask :one
UPDATE tasks
SET name = $1,
    description = $2,
    status = $3,
    updated_at = $4
WHERE id = $5
RETURNING (
        id,
        name,
        description,
        status,
        created_at,
        updated_at
    );
-- name DeleteTask :one
DELETE FROM tasks
WHERE id = $1
RETURNING (
        id,
        name,
        description,
        status,
        created_at,
        updated_at
    );