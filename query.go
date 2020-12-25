package sqlbuilder

import "strings"

// Query describes an sql query.
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
	return &Query{
		str:   &strings.Builder{},
		table: table,
	}
}

func (q *Query) addColumns(columns ...string) {
	for i, c := range columns {
		q.str.WriteString(c)
		if i != len(columns)-1 {
			q.str.WriteByte(',')
		}
	}
}

// Select returns sql select statement.
func (q *Query) Select(columns ...string) *Statement {
	q.reset()
	q.str.WriteString("SELECT ")
	if columns != nil {
		q.addColumns(columns...)
	} else {
		q.str.WriteByte('*')
	}
	q.str.WriteString(" FROM " + q.table)
	return &Statement{q}
}
