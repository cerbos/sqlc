// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
)

const getAll = `-- name: GetAll :many
SELECT id, first_name, last_name, age FROM users
`

func (q *Queries) GetAll(ctx context.Context, db DBTX) ([]User, error) {
	rows, err := db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Age,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}