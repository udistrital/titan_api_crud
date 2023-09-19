package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CampoResolucion_20230918_165217 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CampoResolucion_20230918_165217{}
	m.Created = "20230918_165217"

	migration.Register("CampoResolucion_20230918_165217", m)
}

// Run the migrations
func (m *CampoResolucion_20230918_165217) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230918_165217_campo_resolucion_up.sql")

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
func (m *CampoResolucion_20230918_165217) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230918_165217_campo_resolucion_down.sql")

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
