package models

import "github.com/astaxie/beego/orm"

// TrConceptosNomPersona Estructura para el registro de la transacción, es un arreglo de apuntadores de tipo ConceptoNominaPorPersona
type TrConceptosNomPersona struct {
	Conceptos []*ConceptoNominaPorPersona
}

// RegistrarConceptos transacción para registro de varios conceptos en la tabla concepto_nomina_por_persona
func RegistrarConceptos(m *TrConceptosNomPersona) (alerta Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	_, err = o.InsertMulti(len(m.Conceptos), m.Conceptos)
	if err == nil {
		alerta = Alert{Type: "success", Code: "Ok", Body: err}
		o.Commit()
	} else {
		alerta = Alert{Type: "error", Code: "E_CONCEPTOS", Body: err}
		o.Rollback()
	}
	return alerta, err
}
