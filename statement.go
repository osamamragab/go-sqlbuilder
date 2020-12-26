package sqlbuilder

import "strings"

// Statement describes an sql query statement.
type Statement struct {
	*Query
}

// Where adds sql where condition to query.
// Example:
//	Where("id = $@ AND name = $@", 1, "David")
// "$@" are replaced with argument number.
func (s *Statement) Where(str string, args ...interface{}) *Statement {
	s.str.WriteString(" WHERE ")
	if idx := strings.Index(str, "$@"); idx != -1 {
		var i, last int
		for idx != -1 && i < len(args) {
			s.str.WriteString(str[last : last+idx])
			s.addArg(args[i])
			i++
			last += idx + 2
			idx = strings.Index(str[last:], "$@")
		}
		if len(str) > last {
			s.str.WriteString(str[last:])
		}
	} else {
		s.str.WriteString(str)
	}
	return s
}
