-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1 LIMIT 1;

-- name: GetTodos :many
SELECT * FROM todos ORDER BY id ASC;

-- name: CreateTodo :one
INSERT INTO
    todos (
        title,
        description,
        completed,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: UpdateTodo :one
UPDATE todos
SET
    title = $1,
    description = $2,
    completed = $3,
    updated_at = $4
WHERE
    id = $5
RETURNING
    *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;