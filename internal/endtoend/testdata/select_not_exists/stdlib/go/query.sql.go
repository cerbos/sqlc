// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package querytest

import (
	"context"
)

const barNotExists = `-- name: BarNotExists :one
SELECT
    NOT EXISTS (
        SELECT
            1
        FROM
            bar
        where
            id = $1
    )
`

func (q *Queries) BarNotExists(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, barNotExists, id)
	var not_exists bool
	err := row.Scan(&not_exists)
	return not_exists, err
}
