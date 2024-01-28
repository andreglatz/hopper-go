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

-- name: UpdateLink :exec
UPDATE links SET
  short = $1,
  original = $2,
  clicks = $3
WHERE id = $4;

-- name: GetLinks :many
SELECT * FROM links
WHERE
  (CASE WHEN @short::text IS NOT NULL THEN short LIKE '%' || @short::text || '%' ELSE TRUE END)
  AND (CASE WHEN @original::text IS NOT NULL THEN original LIKE '%' || @original::text || '%' ELSE TRUE END)
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');


-- name: GetLinksCount :one
SELECT COUNT(*) FROM links
WHERE
  (CASE WHEN @short::text IS NOT NULL THEN short LIKE '%' || @short::text || '%' ELSE TRUE END)
  AND (CASE WHEN @original::text IS NOT NULL THEN original LIKE '%' || @original::text || '%' ELSE TRUE END);
