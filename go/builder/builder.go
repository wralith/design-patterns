package builder

import "strings"

// SQLBuilder is an implementation of builder pattern for simple SQL select statements
type SQLBuilder struct {
	statement string
}

// NewBuilder returns a new SQLBuilder
func NewBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// Select add SELECT clause to the statement and returns same SQLBuilder instance
func (b *SQLBuilder) Select(keys ...string) *SQLBuilder {
	b.statement += "SELECT "
	b.statement += strings.Join(keys, ", ")
	b.statement += " "
	return b
}

// From add FROM clause to the statement and returns same SQLBuilder instance
func (b *SQLBuilder) From(table string) *SQLBuilder {
	b.statement += "FROM "
	b.statement += table
	b.statement += " "
	return b
}

// Where add WHERE clause to the statement and returns same SQLBuilder instance
func (b *SQLBuilder) Where(condition string) *SQLBuilder {
	b.statement += "WHERE "
	b.statement += condition
	b.statement += " "
	return b
}

// Build returns the built SQL statement
func (b *SQLBuilder) Build() string {
	return b.statement[:len(b.statement)-1]
}
