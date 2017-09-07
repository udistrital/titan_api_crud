package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DatosPruebas struct {
	Id 									int                      `orm:"auto;column(id);pk"`
	NumDocumento				int 											`orm:"column(num_documento)"`
	MesPreliq						int 											`orm:"column(mes_preliquidacion)"`
	AnoPreliq						int 											`orm:"column(ano_preliquidacion)"`
	ValorSalario				string										`orm:"column(valor_salario)"`
	ValorReteica				string										`orm:"column(valor_reteica)"`
	ValorEstampillaUD		string										`orm:"column(valor_estampillaud)"`
	ValorProCultura			string										`orm:"column(valor_procultura)"`
	ValorAdultoMayor	  string										`orm:"column(valor_adultomayor)"`
	ValorPrimaTecnica	  string										`orm:"column(valor_prima_tecnica)"`
	ValorPrimaAnt			  string										`orm:"column(valor_prima_ant)"`
	ValorSalud 				  string										`orm:"column(valor_salud)"`
	ValorPension    	  string										`orm:"column(valor_pension)"`
}

func (t *DatosPruebas) TableName() string {
	return "datos_pruebas"
}

func init() {
	orm.RegisterModel(new(DatosPruebas))
}

// AddDatosPruebas insert a new DatosPruebas into database and returns
// last inserted Id on success.
func AddDatosPruebas(m *DatosPruebas) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDatosPruebasById retrieves DatosPruebas by Id. Returns error if
// Id doesn't exist
func GetDatosPruebasById(id int) (v *DatosPruebas, err error) {
	o := orm.NewOrm()
	v = &DatosPruebas{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDatosPruebas retrieves all DatosPruebas matches certain condition. Returns empty list if
// no records exist
func GetAllDatosPruebas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	fmt.Println("datos pruebaaaaaaaaaaaaa")
	o := orm.NewOrm()
	qs := o.QueryTable(new(DatosPruebas)).RelatedSel(5)
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

	var l []DatosPruebas
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

// UpdateDatosPruebas updates DatosPruebas by Id and returns error if
// the record to be updated doesn't exist
func UpdateDatosPruebasById(m *DatosPruebas) (err error) {
	o := orm.NewOrm()
	v := DatosPruebas{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDatosPruebas deletes DatosPruebas by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDatosPruebas(id int) (err error) {
	o := orm.NewOrm()
	v := DatosPruebas{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DatosPruebas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
