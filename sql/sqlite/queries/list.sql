-- name: GetList :one
SELECT * FROM lists
WHERE id = ? LIMIT 1;

-- name: AllLists :many
SELECT * FROM lists
ORDER BY id ASC
LIMIT ? OFFSET ?;

-- name: CreateList :one
INSERT INTO lists (name)
VALUES (?)
RETURNING *;

-- name: RenameList :one
UPDATE lists
SET name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists
WHERE id = ?;

