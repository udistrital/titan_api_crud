package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DetallePreliquidacion struct {
	Id                       int                     `orm:"column(id);pk,auto"`
	ContratoPreliquidacionId *ContratoPreliquidacion `orm:"column(contrato_preliquidacion_id);rel(fk)"`
	ValorCalculado           float64                 `orm:"column(valor_calculado)"`
	DiasLiquidados           float64                 `orm:"column(dias_liquidados)"`
	DiasEspecificos          string                  `orm:"column(dias_especificos);null"`
	TipoPreliquidacionId     int                     `orm:"column(tipo_preliquidacion_id);"`
	ConceptoNominaId         *ConceptoNomina         `orm:"column(concepto_nomina_id);rel(fk)"`
	EstadoDisponibilidadId   int                     `orm:"column(estado_disponibilidad_id);"`
	Activo                   bool                    `orm:"column(activo)"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(timestamp with time zone);auto_now_add"`
}

func (t *DetallePreliquidacion) TableName() string {
	return "detalle_preliquidacion"
}

func init() {
	orm.RegisterModel(new(DetallePreliquidacion))
}

// AddDetallePreliquidacion insert a new DetallePreliquidacion into database and returns
// last inserted Id on success.
func AddDetallePreliquidacion(m *DetallePreliquidacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDetallePreliquidacionById retrieves DetallePreliquidacion by Id. Returns error if
// Id doesn't exist
func GetDetallePreliquidacionById(id int) (v *DetallePreliquidacion, err error) {
	o := orm.NewOrm()
	v = &DetallePreliquidacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDetallePreliquidacion retrieves all DetallePreliquidacion matches certain condition. Returns empty list if
// no records exist
func GetAllDetallePreliquidacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DetallePreliquidacion))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DetallePreliquidacion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDetallePreliquidacion updates DetallePreliquidacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateDetallePreliquidacionById(m *DetallePreliquidacion) (err error) {
	o := orm.NewOrm()
	v := DetallePreliquidacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDetallePreliquidacion deletes DetallePreliquidacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDetallePreliquidacion(id int) (err error) {
	o := orm.NewOrm()
	v := DetallePreliquidacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DetallePreliquidacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
