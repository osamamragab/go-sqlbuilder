package sqlbuilder

import "strings"

// Query describes an SQL query.
type Query struct {
	str   *strings.Builder
	args  []interface{}
	table string
}

// reset resets query string and arguments.
func (q *Query) reset() {
	q.str.Reset()
	q.args = nil
}

// String returns query string.
func (q *Query) String() string {
	return q.str.String()
}

// Args returns query arguments.
func (q *Query) Args() []interface{} {
	return q.args
}

// Table returns table name.
func (q *Query) Table() string {
	return q.table
}

// NewQuery returns new Query with table.
func NewQuery(table string) *Query {
	return &Query{table: table}
}
