package models

import (
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
