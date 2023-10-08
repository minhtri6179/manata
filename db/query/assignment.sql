-- name: GetAssignment :one
SELECT *
FROM assignment
WHERE task_id = $1
    AND user_id = $2;