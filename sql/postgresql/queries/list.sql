-- name: GetList :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: AllLists :many
SELECT * FROM lists
ORDER BY id ASC
LIMIT $1
OFFSET $2;

-- name: CreateList :one
INSERT INTO lists (name)
VALUES ($1)
RETURNING *;

-- name: RenameList :one
UPDATE lists
SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists
WHERE id = $1;

