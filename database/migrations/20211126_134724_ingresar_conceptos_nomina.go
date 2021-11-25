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
func (m *IngresarConceptos_20210929_144847) Down() {
	m.SQL("DELETE FROM liquidador.concepto_nomina;")
}
