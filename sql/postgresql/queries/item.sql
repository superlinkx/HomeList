-- name: GetListItem :one
SELECT * FROM items
WHERE id = $1 LIMIT 1;

-- name: AllItemsFromList :many
SELECT * FROM items
WHERE list_id = $1
ORDER BY sort ASC
LIMIT $2;

-- name: CreateListItem :one
INSERT INTO items (list_id, content, sort)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateListItemText :one
UPDATE items
SET content = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListItemSort :one
UPDATE items
SET sort = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListItemChecked :one
UPDATE items
SET checked = $1
WHERE id = $2
RETURNING *;

-- name: DeleteListItem :exec
DELETE FROM items
WHERE id = $1;
