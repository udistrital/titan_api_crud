package models

import "time"

type Persona struct {
	Id       *InformacionProveedor `orm:"column(id);rel(fk)"`
	EmpDesde time.Time             `orm:"column(emp_desde);type(timestamp with time zone)"`
	EmpHasta time.Time             `orm:"column(emp_hasta);type(timestamp with time zone);null"`
	IdCargo  *Cargo                `orm:"column(id_cargo);rel(fk)"`
	Estado   string                `orm:"column(estado)"`
	Regimen  string                `orm:"column(regimen);null"`
}
