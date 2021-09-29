package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IngresarConceptos_20210929_144847 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IngresarConceptos_20210929_144847{}
	m.Created = "20210929_144847"

	migration.Register("IngresarConceptos_20210929_144847", m)
}

// Run the migrations
func (m *IngresarConceptos_20210929_144847) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../20210929_144847_ingresar_conceptos_up")

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
func (m *IngresarConceptos_20210929_144847) Down() {
	file, err := ioutil.ReadFile("../scripts/20210910_234334_crear_bd_down.sql") //se utiliza el mismo archivo de crear bd ya que usa lo mismo

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
