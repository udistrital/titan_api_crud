package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IngresarConceptosNomina_20211125_143514 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IngresarConceptosNomina_20211125_143514{}
	m.Created = "20211125_143514"

	migration.Register("IngresarConceptosNomina_20211125_143514", m)
}

// Run the migrations
func (m *IngresarConceptosNomina_20211125_143514) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20211126_134724_ingresar_conceptos_nomina_up.sql")

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
func (m *IngresarConceptosNomina_20211125_143514) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM liquidador.concepto_nomina")
}
