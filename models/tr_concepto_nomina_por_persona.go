package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// TrConceptosNomPersona Estructura para el registro de la transacción, es un arreglo de apuntadores de tipo ConceptoNominaPorPersona
type TrConceptosNomPersona struct {
	Conceptos []ConceptoNominaPorPersona
}

// RegistrarConceptos transacción para registro de varios conceptos en la tabla concepto_nomina_por_persona
func RegistrarConceptos(m *TrConceptosNomPersona) (alerta Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	var conceptos []int
	for i := 0; i < len(m.Conceptos); i++ {
		id, err := o.Insert(&m.Conceptos[i])
		if err != nil {
			alerta = Alert{Type: "error", Code: "titan_api_crud_error", Body: err}
			o.Rollback()
		} else {
			conceptos = append(conceptos, int(id))
		}
	}
	alerta = Alert{Type: "success", Code: "Ok", Body: conceptos}
	o.Commit()

	return
}

// RegistrarYActualizarIncapacidad cambia el estado del concepto y registra un nuevo, para la prorroga de la incapacidad
func RegistrarYActualizarIncapacidad(m *TrConceptosNomPersona) (conceptos map[string]int, err error) {
	fmt.Println("id: ", m.Conceptos[0].Id)
	o := orm.NewOrm()
	o.Begin()
	incapacidad := m.Conceptos[0]
	incapacidad.Activo = false
	fmt.Println("incapacidad: ", incapacidad)
	_, err = o.Update(incapacidad, "Activo")
	if err != nil {
		o.Rollback()
		return
	}
	conceptos["IdAntiguo"] = incapacidad.Id
	incapacidad.Activo = true
	incapacidad.Id = 0
	id, err := o.Insert(incapacidad)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	conceptos["IdNuevo"] = int(id)
	return
}
