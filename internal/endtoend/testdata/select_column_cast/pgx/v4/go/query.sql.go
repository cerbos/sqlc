// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
)

const selectColumnCast = `-- name: SelectColumnCast :many
SELECT bar::int FROM foo
`

func (q *Queries) SelectColumnCast(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, selectColumnCast)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var bar int32
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
