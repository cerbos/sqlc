// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package querytest

import ()

type Bar struct {
	ID    int64
	Owner string
}

type Foo struct {
	Barid int64
}
