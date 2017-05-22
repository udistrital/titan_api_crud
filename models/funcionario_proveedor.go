package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Funcionario_x_Proveedor struct {
	Id              int     `orm:"column(id_proveedor)"`
	NombreProveedor string  `orm:"column(nom_proveedor)"`
	NumDocumento    float64 `orm:"column(contratista)"`
	NumeroContrato  string  `orm:"column(numero_contrato)"`
	IdEPS                  int  							`orm:"column(id_eps)"`
	IdARL                  int  							`orm:"column(id_arl)"`
	IdFondoPension         int  							`orm:"column(id_fondo_pension)"`
	IdCajaCompensacion     int  							`orm:"column(id_caja_compensacion)"`
}

type Funcionario_x_Pensionado struct {
	Id              int    `orm:"column(informacion_proveedor)"`
	NombreProveedor string `orm:"column(nom_proveedor)"`
	NumDocumento    int    `orm:"column(num_documento)"`
	NumeroContrato  string  `orm:"column(numero_contrato)"`
}

type Funcionario_x_Beneficiario struct {
	Id              int    `orm:"column(id)"`
	InformacionProveedor int `orm:"column(informacion_proveedor)"`
	InformacionPensionado    int    `orm:"column(informacion_pensionado)"`
}

type Funcionario_x_Sustituto struct {
	Id              int    `orm:"column(informacion_proveedor)"`
	NombreProveedor string `orm:"column(nom_proveedor)"`
	NumDocumento    int    `orm:"column(num_documento)"`
	NumeroContrato  string  `orm:"column(numero_contrato)"`
}


func init() {
	orm.RegisterModel(new(Funcionario_x_Proveedor))
}

// last inserted Id on success.
func GetIdProveedorXFuncionario() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()

	var temp []Funcionario_x_Proveedor
	_, err := o.Raw("SELECT informacionproveedor.id_proveedor, informacionproveedor.nom_proveedor,contratos.contratista,contratos.numero_contrato, personanatural.id_eps, personanatural.id_arl, personanatural.id_fondo_pension, personanatural.id_caja_compensacion FROM agora.informacion_proveedor AS informacionproveedor, argo.contrato_general AS contratos, agora.informacion_persona_natural AS personanatural where contratos.objeto_contrato = 'Funcionario de planta' AND contratos.contratista = informacionproveedor.num_documento AND informacionproveedor.num_documento = personanatural.num_documento_persona" ).QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetIdProveedorXDocente() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()

	var temp []Funcionario_x_Proveedor
	_, err := o.Raw("SELECT informacionproveedor.id_proveedor, contratos.contratista,contratos.numero_contrato,informacionproveedor.nom_proveedor FROM agora.informacion_proveedor AS informacionproveedor, argo.contrato_general AS contratos where contratos.objeto_contrato = 'Docente de planta' AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp
}

func GetIdProveedorXContratista() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()

	var temp []Funcionario_x_Proveedor
	_, err := o.Raw("SELECT informacionproveedor.id_proveedor, contratos.contratista,contratos.numero_contrato,informacionproveedor.nom_proveedor FROM agora.informacion_proveedor AS informacionproveedor, argo.contrato_general AS contratos where contratos.objeto_contrato = 'Contratista' AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp
}

func GetIdPensionado() (arregloIDs []Funcionario_x_Pensionado) {
	o := orm.NewOrm()
	var temp []Funcionario_x_Pensionado
	_, err := o.Raw("SELECT pensionado.informacion_proveedor, informacionproveedor.num_documento, informacionProveedor.nom_proveedor,contratos.numero_contrato,contratos.contratista FROM personal.informacion_pensionado AS pensionado, agora.informacion_proveedor AS informacionProveedor,argo.contrato_general AS contratos WHERE contratos.objeto_contrato = 'Pensionado' AND pensionado.informacion_proveedor = informacionproveedor.id_proveedor AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp
}

func GetIdSustituto() (arregloIDs []Funcionario_x_Proveedor) {
	o := orm.NewOrm()
	var temp []Funcionario_x_Proveedor
_, err := o.Raw("SELECT beneficiario.informacion_proveedor, informacionproveedor.num_documento, informacionProveedor.nom_proveedor,contratos.numero_contrato,contratos.contratista FROM personal.sustituto AS sustituto,personal.beneficiario AS beneficiario,agora.informacion_proveedor AS informacionProveedor,argo.contrato_general AS contratos WHERE contratos.objeto_contrato = 'SustPensionado' AND sustituto.beneficiario = beneficiario.id AND beneficiario.informacion_proveedor = informacionproveedor.id_proveedor AND contratos.contratista = informacionproveedor.num_documento").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}
	return temp
}

func ListaContratos(v *Preliquidacion) (datos []Funcionario_x_Proveedor, err error) {
	o := orm.NewOrm()
	consulta := `select c.id_proveedor ,
								      c.nom_proveedor ,
								      b.contratista ,
								      b.numero_contrato
								      from argo.acta_inicio as a inner join argo.contrato_general as b on a.numero_contrato = b.numero_contrato
														   inner join agora.informacion_proveedor as c on b.contratista = c.num_documento
														   inner join argo.tipo_contrato on argo.tipo_contrato.id = b.tipo_contrato
														   where (argo.tipo_contrato.tipo_contrato = ?)  and ((? between a.fecha_inicio and a.fecha_fin) or ( ? between a.fecha_inicio and a.fecha_fin) ); `

	_, err = o.Raw(consulta, v.Nomina.TipoNomina.Nombre, v.FechaInicio, v.FechaFin).QueryRows(&datos)
	return
}
