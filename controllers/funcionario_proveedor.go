package controllers

import (
	"encoding/json"
	//	"errors"
	//"strconv"
	//"strings"
	"github.com/udistrital/titan_api_crud/models"

	"fmt"

	"github.com/astaxie/beego"
)

// ConceptoController operations for Concepto
type FuncionarioProveedorController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuncionarioProveedorController) URLMapping() {
	c.Mapping("FuncionarioProveedor", c.ConsultarIDProveedor)

}

func (c *FuncionarioProveedorController) ConsultarIDProveedor() {
	fmt.Println("hola a todos")
	var v models.Nomina
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.TipoNomina.Nombre == "CT" {
			if listaContratos, err := models.ListaContratosContratistas(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = listaContratos
			} else {
				c.Data["json"] = err.Error()
				fmt.Println("error : ", err)
			}

		} else if v.TipoNomina.Nombre == "HCS" || v.TipoNomina.Nombre == "HCH" {
				if listaContratos, err := models.ListaContratosDocentesDVE(&v); err == nil {
					c.Ctx.Output.SetStatus(201)
					c.Data["json"] = listaContratos
				} else {
					c.Data["json"] = err.Error()
					fmt.Println("error : ", err)
				}
			}

	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}
