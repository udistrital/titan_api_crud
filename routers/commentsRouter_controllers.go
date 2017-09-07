package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DatosPruebasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud2/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
