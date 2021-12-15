package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20211125_120652 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20211125_120652{}
	m.Created = "20211125_120652"

	migration.Register("CrearSchema_20211125_120652", m)
}

// Run the migrations
func (m *CrearSchema_20211125_120652) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS liquidador;")
	m.SQL("ALTER SCHEMA liquidador OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,liquidador;")
}

// Reverse the migrations
func (m *CrearSchema_20211125_120652) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS liquidador")
}
