-- name: GetLink :one
SELECT * FROM links
WHERE id = $1 LIMIT 1;

-- name: ListLinks :many
SELECT * FROM links
ORDER BY short;

-- name: CreateLink :one
INSERT INTO links (
  short, original
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateLink :exec
UPDATE links SET
  short = $2,
  original = $3
WHERE id = $1
RETURNING *;

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1;
