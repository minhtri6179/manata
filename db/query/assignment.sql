-- name CreateAssignment :one
INSERT INTO assignments (
        task_id,
        user_id,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4)
RETURNING (
        task_id,
        user_id,
        created_at,
        updated_at
    );