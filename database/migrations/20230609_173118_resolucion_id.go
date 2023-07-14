package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ResolucionId_20230609_173118 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ResolucionId_20230609_173118{}
	m.Created = "20230609_173118"

	migration.Register("ResolucionId_20230609_173118", m)
}

// Run the migrations
func (m *ResolucionId_20230609_173118) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230609_173118_resolucion_id_up.sql")

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
func (m *ResolucionId_20230609_173118) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/220230609_173118_resolucion_id_down.sql")

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
