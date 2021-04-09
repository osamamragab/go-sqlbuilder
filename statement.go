package sqlbuilder

// Statement describes an sql query statement.
type Statement struct {
	*Query
}

// Where adds sql where condition to query.
func (s *Statement) Where(cond string, args ...interface{}) *Statement {
	s.str.WriteString(" WHERE ")
	s.Raw(cond, args...)
	return s
}

// Limit adds sql limit to query.
//
// Limit panics if n <= 0.
func (s *Statement) Limit(n int) *Statement {
	if n <= 0 {
		panic("sqlbuilder: invalid limit value")
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
		panic("sqlbuilder: invalid offset value")
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
	if len(columns) > 0 {
		s.str.WriteString(" ORDER BY ")
		s.addColumns(columns...)
		s.str.WriteString(" DESC")
	}
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
