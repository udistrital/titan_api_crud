package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type Preliquidacion struct {
	Nomina               *Nomina               `orm:"column(nomina);rel(fk)"`
	Id 									int                      `orm:"auto;column(id);pk"`
	Descripcion          string                `orm:"column(descripcion);null"`
	Mes                  int                   `orm:"column(mes)"`
	Ano                  int                   `orm:"column(ano)"`
	FechaRegistro        time.Time             `orm:"column(fecha_registro);type(timestamp with time zone)"`
	EstadoPreliquidacion *EstadoPreliquidacion `orm:"column(estado_preliquidacion);rel(fk)"`
}

type InformePreliquidacion struct {
	//IdPersona      int     `orm:"column(id)"`
	NumeroContrato  string
	Vigencia       int
	Conceptos      []ConceptosInforme
	Disponibilidad int
}


type ConceptosInforme struct {
	Id         int    `orm:"column(id)"`
	Nombre     string `orm:"column(nombre)"`
	Naturaleza string `orm:"column(naturaleza)"`
	Valor      string `orm:"column(valor)"`
	TipoPreliquidacion string `orm:"column(tipo)"`
	EstadoDisponibilidad int `orm:"column(id_disp)"`
}

type Preliquidacion_x_contratos struct {
	Id_Preliq int `orm:"column(id)"`
	Nombre_tipo_nomina string  `orm:"column(nombre)"`
	Contratos_por_preliq []Contrato_x_Vigencia
}

type Contrato_x_Vigencia struct {
	NumeroContrato string `orm:"column(numero_contrato)"`
	VigenciaContrato int `orm:"column(vigencia_contrato)"`
}

type Totales_x_preliq struct {
	Total float64    `orm:"column(total);null"`
	Id_concepto int   `orm:"column(id)"`
	Alias_concepto string   `orm:"column(alias_concepto)"`
	Nombre_concepto   string  `orm:"column(nombre_concepto)"`
}

func (t *Preliquidacion) TableName() string {
	return "preliquidacion"
}

func init() {
	orm.RegisterModel(new(Preliquidacion))
}

// AddPreliquidacion insert a new Preliquidacion into database and returns
// last inserted Id on success.
func AddPreliquidacion(m *Preliquidacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPreliquidacionById retrieves Preliquidacion by Id. Returns error if
// Id doesn't exist
func GetPreliquidacionById(id int) (v *Preliquidacion, err error) {
	o := orm.NewOrm()
	v = &Preliquidacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPreliquidacion retrieves all Preliquidacion matches certain condition. Returns empty list if
// no records exist
func GetAllPreliquidacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Preliquidacion)).RelatedSel(5)
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

	var l []Preliquidacion
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

// UpdatePreliquidacion updates Preliquidacion by Id and returns error if
// the record to be updated doesn't exist
func UpdatePreliquidacionById(m *Preliquidacion) (err error) {
	o := orm.NewOrm()
	v := Preliquidacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePreliquidacion deletes Preliquidacion by Id and returns error if
// the record to be deleted doesn't exist
func DeletePreliquidacion(id int) (err error) {
	o := orm.NewOrm()
	v := Preliquidacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Preliquidacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ResumenPreliquidacion(v *Preliquidacion) (resumen []InformePreliquidacion, err error) {
	o := orm.NewOrm()
	var numero_contratos []Contrato_x_Vigencia
	var est_disp int

	_, err = o.Raw("select numero_contrato, vigencia_contrato from administrativa.detalle_preliquidacion where preliquidacion = ? group by numero_contrato,vigencia_contrato", v.Id).QueryRows(&numero_contratos)
	if numero_contratos != nil && err == nil {

		for _, contrato := range numero_contratos {

			var informe InformePreliquidacion
			informe.NumeroContrato = contrato.NumeroContrato
			informe.Vigencia = contrato.VigenciaContrato
				_, err = o.Raw("SELECT concepto.id as id, concepto.alias_concepto as nombre, naturaleza.nombre as naturaleza, detalle.valor_calculado as valor, detalle.estado_disponibilidad as id_disp, tipo.nombre as tipo from administrativa.detalle_preliquidacion as detalle, administrativa.concepto_nomina as concepto, administrativa.naturaleza_concepto_nomina as naturaleza, administrativa.tipo_preliquidacion as tipo WHERE detalle.concepto = concepto.id AND concepto.naturaleza_concepto = naturaleza.id AND detalle.tipo_preliquidacion = tipo.id AND detalle.preliquidacion = ? AND detalle.numero_contrato = ? AND detalle.vigencia_contrato = ?",v.Id, contrato.NumeroContrato,contrato.VigenciaContrato).QueryRows(&informe.Conceptos)
				if err != nil {
					fmt.Println("err3: ", err)
				}
			for _, concepto := range informe.Conceptos {
			   est_disp = 2
				if(concepto.EstadoDisponibilidad == 1){

					est_disp = 1
				}
			}

			informe.Disponibilidad = est_disp

			resumen = append(resumen, informe)

		}


	} else {
		fmt.Println("err1: ", err)
	}
	return
}



func Contratos_x_preliquidacion(idNomina, mes, ano int) (cont_por_pre Preliquidacion_x_contratos, err error) {
	o := orm.NewOrm()

	var preliq_x_cont  Preliquidacion_x_contratos

	_ = o.Raw("select tipo_nom.nombre, pre.id from administrativa.preliquidacion as pre, administrativa.nomina as nom, administrativa.tipo_nomina as tipo_nom where pre.ano = ? AND pre.mes=? AND pre.estado_preliquidacion = 4 AND nomina = ? AND pre.nomina = nom.id AND nom.tipo_nomina = tipo_nom.id;", ano, mes, idNomina).QueryRow(&preliq_x_cont.Nombre_tipo_nomina,&preliq_x_cont.Id_Preliq)
	if err == nil {

		_, err = o.Raw("select detalle.numero_contrato, detalle.vigencia_contrato from administrativa.detalle_preliquidacion as detalle, administrativa.preliquidacion as pre where detalle.preliquidacion = pre.id AND pre.ano = ? AND pre.mes=? AND pre.estado_preliquidacion = 4 AND nomina = ?  group by detalle.numero_contrato,detalle.vigencia_contrato;", ano, mes, idNomina).QueryRows(&preliq_x_cont.Contratos_por_preliq)
		if err == nil {
			fmt.Println(preliq_x_cont)

		} else {
			fmt.Println("err1: ", err)
		}
	}else{
		fmt.Println("err1: ", err)
	}
	return preliq_x_cont,err
}

func Totales_ss_x_preliquidacion(idNomina, mes, ano int) (totales_por_pre []Totales_x_preliq, err error) {
	o := orm.NewOrm()

	var totales  []Totales_x_preliq

	_,err = o.Raw("SELECT SUM(valor_calculado) as total, cn.alias_concepto, cn.nombre_concepto, cn.id FROM administrativa.detalle_preliquidacion dp INNER JOIN administrativa.concepto_nomina cn ON cn.id = dp.concepto INNER JOIN  administrativa.preliquidacion pr ON pr.id = dp.preliquidacion WHERE cn.nombre_concepto = 'salud' OR cn.nombre_concepto = 'pension' OR cn.nombre_concepto = 'fondoSolidaridad' AND pr.mes = ? AND pr.ano = ? AND pr.estado_preliquidacion = 1 AND pr.nomina = ? GROUP BY cn.id, cn.alias_concepto, cn.nombre_concepto, cn.id;", mes, ano, idNomina).QueryRows(&totales)
	if err == nil {
		//fmt.Println(totales)

	}else{
		fmt.Println("err1: ", err)
	}
	return totales,err
}

func Contratos_x_preliquidacion_cerrada(idNomina, mes, ano int) (cont_por_pre Preliquidacion_x_contratos, err error) {
	o := orm.NewOrm()
	var preliq_x_cont  Preliquidacion_x_contratos

	err1 := o.Raw("select tipo_nom.nombre, pre.id from administrativa.preliquidacion as pre, administrativa.nomina as nom, administrativa.tipo_nomina as tipo_nom where pre.ano = ? AND pre.mes=? AND pre.estado_preliquidacion = 1 AND nomina = ? AND pre.nomina = nom.id AND nom.tipo_nomina = tipo_nom.id;", ano, mes, idNomina).QueryRow(&preliq_x_cont.Nombre_tipo_nomina,&preliq_x_cont.Id_Preliq)
	if err1 == nil {

		_, err = o.Raw("select detalle.numero_contrato, detalle.vigencia_contrato from administrativa.detalle_preliquidacion as detalle, administrativa.preliquidacion as pre where detalle.preliquidacion = pre.id AND pre.ano = ? AND pre.mes=? AND pre.estado_preliquidacion = 4 AND nomina = ?  group by detalle.numero_contrato,detalle.vigencia_contrato;", ano, mes, idNomina).QueryRows(&preliq_x_cont.Contratos_por_preliq)
		if err == nil {
			fmt.Println(preliq_x_cont)

		} else {
			fmt.Println("err1: ", err)
		}
	}else{
		fmt.Println("err1: ", err1)
	}

	return preliq_x_cont,err1
}
