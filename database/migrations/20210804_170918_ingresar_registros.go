package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IngresarRegistros_20210804_170918 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IngresarRegistros_20210804_170918{}
	m.Created = "20210804_170918"

	migration.Register("IngresarRegistros_20210804_170918", m)
}

// Run the migrations
func (m *IngresarRegistros_20210804_170918) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *IngresarRegistros_20210804_170918) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
