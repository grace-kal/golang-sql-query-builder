package querybuilder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	tableName string
	query     string
}

// The NewQueryBuilder function creates a new instance of QueryBuilder
func NewQueryBuilder(tableName string) *QueryBuilder {
	return &QueryBuilder{
		tableName: tableName,
	}
}

// The Create function generates a SQL CREATE TABLE query
// Entry parameter is the set of key-value items namely the column names and their data types used to CREATE TABLE
func (qb *QueryBuilder) Create(data map[string]string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("CREATE TABLE %s (", qb.tableName))
	for columnName, columnType := range data {
		sb.WriteString(fmt.Sprintf("%s %s, ", columnName, columnType))
	}
	sb.WriteString(")")
	return sb.String()
}

// The Select function generates a SQL SELECT query
// Entry parameter is the array of column names to use in the SELECT query
func (qb *QueryBuilder) Select(columns []string) string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(fmt.Sprintf(" FROM %s", qb.tableName))
	return sb.String()
}

// The Insert function generates a SQL INSERT query
// Entry parameter is the set of key-value items to INSERT in the QueryBuilder initialized table name
func (qb *QueryBuilder) Insert(data map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("INSERT INTO %s (", qb.tableName))
	var columns []string
	var values []string
	for col, val := range data {
		columns = append(columns, col)
		values = append(values, fmt.Sprintf("'%v'", val))
	}
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") VALUES (")
	sb.WriteString(strings.Join(values, ", "))
	sb.WriteString(")")
	return sb.String()
}

// The Update function generates a SQL UPDATE query
// Entry parameter is the set of key-value items to UPDATE in the QueryBuilder initialized table name
func (qb *QueryBuilder) Update(data map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("UPDATE %s SET ", qb.tableName))
	var updates []string
	for col, val := range data {
		updates = append(updates, fmt.Sprintf("%s = '%v'", col, val))
	}
	sb.WriteString(strings.Join(updates, ", "))
	return sb.String()
}
