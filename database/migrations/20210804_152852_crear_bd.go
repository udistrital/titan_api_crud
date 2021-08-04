package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearBd_20210804_152852 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearBd_20210804_152852{}
	m.Created = "20210804_152852"

	migration.Register("CrearBd_20210804_152852", m)
}

// Run the migrations
func (m *CrearBd_20210804_152852) Up() {
	file, err := ioutil.ReadFile("../scripts/20210804_152852_crear_bd_up.sql")

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
func (m *CrearBd_20210804_152852) Down() {
	file, err := ioutil.ReadFile("../scripts/20210804_152852_crear_bd_down.sql")

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
