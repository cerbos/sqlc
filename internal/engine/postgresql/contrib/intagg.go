// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

var funcsIntagg = []*catalog.Function{
	{
		Name: "int_array_aggregate",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer[]"},
	},
	{
		Name: "int_array_enum",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
}

func Intagg() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsIntagg
	return s
}
