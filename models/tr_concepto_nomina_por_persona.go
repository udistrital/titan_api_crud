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

// RegistrarYActualizarIncapacidad cambia el estado del concepto y registra un nuevo, para la prorroga de la incapacidad
func RegistrarYActualizarIncapacidad(m *TrConceptosNomPersona) (map[string]interface{}, error) {
	o := orm.NewOrm()
	o.Begin()
	conceptos := make(map[string]interface{})
	incapacidad := m.Conceptos[0]
	incapacidad.Activo = false

	_, err := o.Update(&incapacidad, "activo")
	if err != nil {
		o.Rollback()
		return conceptos, err
	}

	conceptos["Id"] = incapacidad.Id
	incapacidad.Activo = true
	incapacidad.Id = 0
	id, err := o.Insert(&incapacidad)
	if err != nil {
		o.Rollback()
		return conceptos, err
	}
	o.Commit()
	conceptos["IdProrroga"] = int(id)
	return conceptos, err
}

// EliminarIncapacidad esta función se llama en caso de que al registrar una prorroga de incapacidad, el api de ss este abajo
// toma el registro hecho para la prorroga, le cambia el estado a inactivo, y luego tomar la incapacidad anterior y vuleve
// su estado activo
func EliminarIncapacidad(m []map[string]interface{}) error {
	o := orm.NewOrm()
	o.Begin()
	infoRegistros := m[0]

	incapacidadProrroga := ConceptoNominaPorPersona{Id: int(infoRegistros["IdProrroga"].(float64)), Activo: false}
	incapacidadInactiva := ConceptoNominaPorPersona{Id: int(infoRegistros["Id"].(float64)), Activo: true}

	_, err := o.Update(&incapacidadProrroga, "activo")
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.Update(&incapacidadInactiva, "activo")
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return err
}
