package sqlbuilder

import (
	"strconv"
	"strings"
)

// Query describes an sql query.
type Query struct {
	str   *strings.Builder
	args  []interface{}
	table string
}

// NewQuery returns new Query with table.
func NewQuery(table string) *Query {
	return &Query{
		str:   &strings.Builder{},
		table: table,
	}
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

func (q *Query) addColumns(columns ...string) {
	for i, c := range columns {
		q.str.WriteString(c)
		if i != len(columns)-1 {
			q.str.WriteByte(',')
		}
	}
}

func (q *Query) addArg(arg interface{}) {
	q.args = append(q.args, arg)
	q.str.WriteString("$" + strconv.Itoa(len(q.args)))
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

// Insert returns sql insert statement.
func (q *Query) Insert(columns []string, values ...[]string) *Statement {
	q.reset()
	q.str.WriteString("INSERT INTO " + q.table + " (")
	q.addColumns(columns...)
	q.str.WriteString(") VALUES ")
	for i, vs := range values {
		q.str.WriteByte('(')
		for j, v := range vs {
			q.addArg(v)
			if j != len(vs)-1 {
				q.str.WriteByte(',')
			}
		}
		q.str.WriteByte(')')
		if i != len(values)-1 {
			q.str.WriteByte(',')
		}
	}
	return &Statement{q}
}
