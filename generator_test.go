package main

import (
	"strings"
	"testing"
)

func TestGeneratorList(t *testing.T) {
	generator := Generator{}

	table := createTestTable()
	sql := generator.list(table)

	expected := `select id, first_name, last_name, foo from generator_test`
	if !strings.EqualFold(sql, expected) {
		t.Errorf("Unexpected result! %s != %s", sql, expected)
	}
}

func TestGeneratorByID(t *testing.T) {
	generator := Generator{}

	table := createTestTable()
	sql := generator.byID(table)

	expected := `select id, first_name, last_name, foo from generator_test where id = ?`
	if !strings.EqualFold(sql, expected) {
		t.Errorf("Unexpected result! %s != %s", sql, expected)
	}
}

func createTestTable() *Table {
	table := new(Table)

	table.name = "generator_test"

	table.columns = append(table.columns, &Column{name: "id", ctype: "int(11)", null: "NO", key: "PRI", extra: "auto_increment"})
	table.columns = append(table.columns, &Column{name: "first_name", ctype: "varchar(32)", null: "NO"})
	table.columns = append(table.columns, &Column{name: "last_name", ctype: "varchar(32)", null: "YES"})
	table.columns = append(table.columns, &Column{name: "foo", ctype: "varchar(32)", null: "NO"})

	table.columnNames = append(table.columnNames, "id")
	table.columnNames = append(table.columnNames, "first_name")
	table.columnNames = append(table.columnNames, "last_name")
	table.columnNames = append(table.columnNames, "foo")

	return table
}
