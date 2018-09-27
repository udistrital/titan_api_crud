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
	var registros []int64
	fmt.Println("concepto 0: ", m.Conceptos[0])
	o := orm.NewOrm()
	o.Begin()

	for i := 0; i < len(m.Conceptos); i++ {
		id, err := o.Insert(&m.Conceptos[i])
		if err != nil {
			alerta = Alert{Type: "error", Code: "E_CONCEPTOS", Body: err}
			o.Rollback()
		} else {
			registros = append(registros, id)
		}
	}

	alerta = Alert{Type: "success", Code: "Ok", Body: registros}
	o.Commit()

	return
}
