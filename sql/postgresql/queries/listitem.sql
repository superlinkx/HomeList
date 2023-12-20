-- name: GetListItem :one
SELECT * FROM list_items
WHERE id = $1 LIMIT 1;

-- name: AllItemsFromList :many
SELECT * FROM list_items
WHERE list_id = $1
ORDER BY sort ASC
LIMIT $2;

-- name: CreateListItem :one
INSERT INTO list_items (list_id, content, sort)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateListItemText :one
UPDATE list_items
SET content = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListItemSort :one
UPDATE list_items
SET sort = $1
WHERE id = $2
RETURNING *;

-- name: UpdateListItemChecked :one
UPDATE list_items
SET checked = $1
WHERE id = $2
RETURNING *;

-- name: DeleteListItem :exec
DELETE FROM list_items
WHERE id = $1;
