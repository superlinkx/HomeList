-- name: GetListItem :one
SELECT * FROM list_items
WHERE id = ? LIMIT 1;

-- name: AllItemsFromList :many
SELECT * FROM list_items
WHERE list_id = ?
ORDER BY sort ASC
LIMIT ?;

-- name: CreateListItem :one
INSERT INTO list_items (list_id, content, sort)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateListItemText :one
UPDATE list_items
SET content = ?
WHERE id = ?
RETURNING *;

-- name: UpdateListItemSort :one
UPDATE list_items
SET sort = ?
WHERE id = ?
RETURNING *;

-- name: DeleteListItem :exec
DELETE FROM list_items
WHERE id = ?;
