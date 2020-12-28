package sqlbuilder

import (
	"strconv"
	"strings"
)

// Statement describes an sql query statement.
type Statement struct {
	*Query
}

// Where adds sql where condition to query.
func (s *Statement) Where(str string, args ...interface{}) *Statement {
	s.str.WriteString(" WHERE ")

	if s.driver == "pg" {
		idx := strings.IndexByte(str, '?')
		if idx != -1 {
			var i, last int
			for idx != -1 && i < len(args) {
				s.str.WriteString(str[last : last+idx])
				s.args = append(s.args, args[i])
				s.str.WriteString("$" + strconv.Itoa(len(s.args)))

				i++
				last += idx + 2
				idx = strings.IndexByte(str[last:], '?')
			}
			if len(str) > last {
				s.str.WriteString(str[last:])
			}
			return s
		}
	}

	s.str.WriteString(str)
	if args != nil {
		s.args = append(s.args, args...)
	}

	return s
}

// Limit adds sql limit to query.
//
// Limit panics if n <= 0.
func (s *Statement) Limit(n int) *Statement {
	if n <= 0 {
		panic("invalid limit value")
	}
	s.str.WriteString(" LIMIT ")
	s.addArg(n)
	return s
}

// Offset adds sql offset to query.
//
// Offset panics if n <= 0.
func (s *Statement) Offset(n int) *Statement {
	if n <= 0 {
		panic("invalid offset value")
	}
	s.str.WriteString(" OFFSET ")
	s.addArg(n)
	return s
}

// OrderBy adds sql order by columns asc to query.
func (s *Statement) OrderBy(columns ...string) *Statement {
	if len(columns) > 0 {
		s.str.WriteString(" ORDER BY ")
		s.addColumns(columns...)
	}
	return s
}

// OrderByDesc adds sql order by columns desc to query.
func (s *Statement) OrderByDesc(columns ...string) *Statement {
	s.OrderBy(columns...)
	s.str.WriteString(" DESC")
	return s
}

// Returning adds sql returning to query.
// Should be used with insert or update.
func (s *Statement) Returning(columns ...string) *Statement {
	if len(columns) > 0 {
		s.str.WriteString(" RETURNING ")
		s.addColumns(columns...)
	}
	return s
}
