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


	} else {
		c.Data["json"] = err.Error()
		fmt.Println("error 2: ", err)
	}
	c.ServeJSON()
}

// GetIdProveedorXFuncionario ...
// @Title GetIdProveedorXFuncionario
// @Description Retorna los contratos de planta
// @Success 201 {object} models.Funcionario_x_Proveedor
// @Failure 403 body is empty
// @router /get_funcionarios_planta [get]
func (c *FuncionarioProveedorController) GetIdProveedorXFuncionario(){

	respuesta := models.GetIdProveedorXFuncionario();

	if(respuesta != nil){
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
	}else{
		c.Data["json"] = respuesta

	}
	c.ServeJSON()
}
