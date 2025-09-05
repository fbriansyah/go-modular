package database

import "fmt"

// QueryBuilder provides a fluent interface for building SQL queries
type QueryBuilder struct {
	query string
	args  []interface{}
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		args: make([]interface{}, 0),
	}
}

// Select starts a SELECT query
func (qb *QueryBuilder) Select(columns string) *QueryBuilder {
	qb.query = fmt.Sprintf("SELECT %s", columns)
	return qb
}

// From adds a FROM clause
func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.query += fmt.Sprintf(" FROM %s", table)
	return qb
}

// Where adds a WHERE clause
func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.query += fmt.Sprintf(" WHERE %s", condition)
	qb.args = append(qb.args, args...)
	return qb
}

// And adds an AND condition
func (qb *QueryBuilder) And(condition string, args ...interface{}) *QueryBuilder {
	qb.query += fmt.Sprintf(" AND %s", condition)
	qb.args = append(qb.args, args...)
	return qb
}

// Or adds an OR condition
func (qb *QueryBuilder) Or(condition string, args ...interface{}) *QueryBuilder {
	qb.query += fmt.Sprintf(" OR %s", condition)
	qb.args = append(qb.args, args...)
	return qb
}

// OrderBy adds an ORDER BY clause
func (qb *QueryBuilder) OrderBy(column, direction string) *QueryBuilder {
	qb.query += fmt.Sprintf(" ORDER BY %s %s", column, direction)
	return qb
}

// Limit adds a LIMIT clause
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.query += " LIMIT $" + fmt.Sprintf("%d", len(qb.args)+1)
	qb.args = append(qb.args, limit)
	return qb
}

// Offset adds an OFFSET clause
func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.query += " OFFSET $" + fmt.Sprintf("%d", len(qb.args)+1)
	qb.args = append(qb.args, offset)
	return qb
}

// Build returns the final query and arguments
func (qb *QueryBuilder) Build() (string, []interface{}) {
	return qb.query, qb.args
}

// String returns the query as a string
func (qb *QueryBuilder) String() string {
	return qb.query
}
