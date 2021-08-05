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
	file, err := ioutil.ReadFile("../scripts/20210804_170918_ingresar_registros_up.sql")

	if err != nil {
	// handle error
	fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
	fmt.Println(request)
	m.SQL(request)
	// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *IngresarRegistros_20210804_170918) Down() {
	file, err := ioutil.ReadFile("../scripts/20210804_170918_ingresar_registros_down.sql")

	if err != nil {
	// handle error
	fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
	fmt.Println(request)
	m.SQL(request)
	// do whatever you need with result and error
	}
}
