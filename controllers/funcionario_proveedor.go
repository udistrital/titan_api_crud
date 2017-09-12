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
	var v models.Nomina
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.TipoNomina.Nombre == "HCH" || v.TipoNomina.Nombre == "HCS" {
			if listaContratos, err := models.ListaContratos(&v); err == nil {
				c.Ctx.Output.SetStatus(201)

				c.Data["json"] = listaContratos
			} else {
				c.Data["json"] = err.Error()
				fmt.Println("error : ", err)
			}
		} else if v.TipoNomina.Nombre == "FP" {
			listaContratos := models.GetIdProveedorXFuncionario()
			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = listaContratos

		} else if v.TipoNomina.Nombre == "DP" {
			listaContratos := models.GetIdProveedorXDocente()
			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = listaContratos

		} else if v.TipoNomina.Nombre == "PE" {
			listaContratos := models.GetIdPensionado()
			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = listaContratos

		} else if v.TipoNomina.Nombre == "CT" {
			listaContratos := models.GetIdProveedorXContratista()
			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = listaContratos

		}
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}
