-- name: GetLinkByShort :one
SELECT * FROM links
WHERE short = $1 LIMIT 1;

-- name: CreateLink :one
INSERT INTO links (
  short, original
) VALUES (
  $1, $2
)
RETURNING *;
