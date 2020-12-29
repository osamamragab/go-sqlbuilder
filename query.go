package sqlbuilder

import (
	"reflect"
	"strconv"
	"strings"
)

// Query describes an sql query.
type Query struct {
	str    *strings.Builder
	args   []interface{}
	table  string
	driver string
}

// NewQuery returns new Query with table.
func NewQuery(table string) *Query {
	return &Query{
		str:    &strings.Builder{},
		table:  table,
		driver: "pg",
	}
}

// SetDriver sets driver field to the given value.
// SetDriver panics if driver is not supported.
func (q *Query) SetDriver(driver string) {
	switch d := strings.ToLower(driver); d {
	case "pg", "postgres", "postgresql":
		q.driver = "pg"
	case "mysql":
		q.driver = d
	default:
		panic("unsupported driver: " + driver)
	}
}

// Reset resets query string and arguments.
func (q *Query) Reset() {
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

// SetTable sets table field and calls Reset.
func (q *Query) SetTable(table string) {
	q.Reset()
	q.table = table
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
	switch q.driver {
	case "pg":
		q.str.WriteString("$" + strconv.Itoa(len(q.args)))
	case "mysql":
		q.str.WriteByte('?')
	}
}

// Select returns sql select statement.
func (q *Query) Select(columns ...string) *Statement {
	q.Reset()
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
func (q *Query) Insert(columns []string, values ...interface{}) *Statement {
	q.Reset()
	q.str.WriteString("INSERT INTO " + q.table + " (")
	q.addColumns(columns...)
	q.str.WriteString(") VALUES (")
	var multiple bool
	for i, vs := range values {
		v := reflect.ValueOf(vs)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			if !multiple {
				multiple = true
			}
			if i != 0 {
				q.str.WriteByte('(')
			}
			for j := 0; j < v.Len(); j++ {
				q.addArg(v.Index(j).Interface())
				if j != v.Len()-1 {
					q.str.WriteByte(',')
				}
			}
			q.str.WriteByte(')')
			if i != len(values)-1 {
				q.str.WriteByte(',')
			}
		} else {
			if multiple {
				panic("invalid values type")
			}
			q.addArg(vs)
			if i != len(values)-1 {
				q.str.WriteByte(',')
			}
		}
	}
	if !multiple {
		q.str.WriteByte(')')
	}
	return &Statement{q}
}

// Update returns sql update statement.
// data type could be string or map[string]interface{}.
func (q *Query) Update(data interface{}, args ...interface{}) *Statement {
	q.Reset()
	q.str.WriteString("UPDATE " + q.table + " SET ")
	switch d := data.(type) {
	case string:
		q.str.WriteString(d)
	case map[string]interface{}:
		i := len(d) - 1
		for k, v := range d {
			q.str.WriteString(k + "=")
			q.addArg(v)
			if i != 0 {
				q.str.WriteByte(',')
			}
			i--
		}
	default:
		panic("unexpected data type")
	}
	if args != nil {
		q.args = append(q.args, args...)
	}
	return &Statement{q}
}

// Delete returns sql delete statement.
func (q *Query) Delete() *Statement {
	q.Reset()
	q.str.WriteString("DELETE FROM " + q.table)
	return &Statement{q}
}

// Raw wirtes raw string to query and appends args to query arguments.
func (q *Query) Raw(str string, args ...interface{}) {
	q.str.WriteString(str)
	q.args = append(q.args, args...)
}
