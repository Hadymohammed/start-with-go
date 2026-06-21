-- name: GetNote :one
SELECT * FROM notes WHERE id = $1;

-- name: ListNotes :many
SELECT * FROM notes ORDER BY created_at DESC;

-- name: CreateNote :one
INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING *;

-- name: UpdateNote :one
UPDATE notes SET title = $1, content = $2, updated_at = NOW() WHERE id = $3 RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;
