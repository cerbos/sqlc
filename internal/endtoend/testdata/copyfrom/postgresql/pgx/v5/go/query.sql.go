// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type InsertValuesParams struct {
	A pgtype.Text
	B pgtype.Int4
}