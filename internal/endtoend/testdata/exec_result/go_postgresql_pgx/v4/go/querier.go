// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import (
	"context"

	"github.com/jackc/pgconn"
)

type Querier interface {
	DeleteBarByID(ctx context.Context, id int32) (pgconn.CommandTag, error)
}

var _ Querier = (*Queries)(nil)
