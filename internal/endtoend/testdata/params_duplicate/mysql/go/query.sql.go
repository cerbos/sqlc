// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const selectUserByID = `-- name: SelectUserByID :many
SELECT first_name from
users where (? = id OR ? = 0)
`

type SelectUserByIDParams struct {
	ID int32
}

func (q *Queries) SelectUserByID(ctx context.Context, arg SelectUserByIDParams) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserByID, arg.ID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserByName = `-- name: SelectUserByName :many
SELECT first_name
FROM users
WHERE first_name = ?
   OR last_name = ?
`

type SelectUserByNameParams struct {
	Name sql.NullString
}

func (q *Queries) SelectUserByName(ctx context.Context, arg SelectUserByNameParams) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserByName, arg.Name, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectUserQuestion = `-- name: SelectUserQuestion :many
SELECT first_name from
users where (? = id OR  ? = 0)
`

type SelectUserQuestionParams struct {
	ID      int32
	Column2 interface{}
}

func (q *Queries) SelectUserQuestion(ctx context.Context, arg SelectUserQuestionParams) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, selectUserQuestion, arg.ID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var first_name sql.NullString
		if err := rows.Scan(&first_name); err != nil {
			return nil, err
		}
		items = append(items, first_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
