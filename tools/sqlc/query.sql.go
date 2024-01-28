// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sql

import (
	"context"
)

const createLink = `-- name: CreateLink :one
INSERT INTO links (
  short, original
) VALUES (
  $1, $2
)
RETURNING id, short, original
`

type CreateLinkParams struct {
	Short    string
	Original string
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.db.QueryRow(ctx, createLink, arg.Short, arg.Original)
	var i Link
	err := row.Scan(&i.ID, &i.Short, &i.Original)
	return i, err
}

const deleteLink = `-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1
`

func (q *Queries) DeleteLink(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteLink, id)
	return err
}

const getLink = `-- name: GetLink :one
SELECT id, short, original FROM links
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, id int32) (Link, error) {
	row := q.db.QueryRow(ctx, getLink, id)
	var i Link
	err := row.Scan(&i.ID, &i.Short, &i.Original)
	return i, err
}

const listLinks = `-- name: ListLinks :many
SELECT id, short, original FROM links
ORDER BY short
`

func (q *Queries) ListLinks(ctx context.Context) ([]Link, error) {
	rows, err := q.db.Query(ctx, listLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(&i.ID, &i.Short, &i.Original); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLink = `-- name: UpdateLink :exec
UPDATE links SET
  short = $2,
  original = $3
WHERE id = $1
RETURNING id, short, original
`

type UpdateLinkParams struct {
	ID       int32
	Short    string
	Original string
}

func (q *Queries) UpdateLink(ctx context.Context, arg UpdateLinkParams) error {
	_, err := q.db.Exec(ctx, updateLink, arg.ID, arg.Short, arg.Original)
	return err
}