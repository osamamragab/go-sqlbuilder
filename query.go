package sqlbuilder

import (
	"strings"
)

// Query describes an SQL query.
type Query struct {
	str   *strings.Builder
	args  []interface{}
	table string
}

// NewQuery returns new Query with table.
func NewQuery(table string) *Query {
	return &Query{table: table}
}
