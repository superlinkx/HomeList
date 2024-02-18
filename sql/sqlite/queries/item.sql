-- name: GetListItem :one
SELECT * FROM items
WHERE id = ? LIMIT 1;

-- name: AllItemsFromList :many
SELECT * FROM items
WHERE list_id = ?
ORDER BY sort ASC
LIMIT ?;

-- name: CreateListItem :one
INSERT INTO items (list_id, content, sort)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateListItemText :one
UPDATE items
SET content = ?
WHERE id = ?
RETURNING *;

-- name: UpdateListItemSort :one
UPDATE items
SET sort = ?
WHERE id = ?
RETURNING *;

-- name: UpdateListItemChecked :one
UPDATE items
SET checked =?
WHERE id =?
RETURNING *;

-- name: DeleteListItem :exec
DELETE FROM items
WHERE id = ?;
