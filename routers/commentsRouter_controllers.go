package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ActaInicioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ArgoOrdenadoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CargoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CategoriaBeneficiarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaPorPersonaController"],
		beego.ControllerComments{
			Method: "TrConceptosPorPersona",
			Router: `/TrConceptosPorPersona`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
		beego.ControllerComments{
			Method: "GetPersonasPagosPendientes",
			Router: `/get_personas_pago_pendiente`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
		beego.ControllerComments{
			Method: "ConsultarAsignacionBasica",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"],
			beego.ControllerComments{
				Method: "ConsultarAsignacionBasicaFuncionario",
				Router: `/get_asignacion_basica`,
				AllowHTTPMethods: []string{"post"},
				MethodParams: param.Make(),
				Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
		beego.ControllerComments{
			Method: "ConsultarCedulaDocente",
			Router: `/consultarCedulaDocente/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoCivilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:EstadoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"],
		beego.ControllerComments{
			Method: "ConsultarPrimaTecnica",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"],
		beego.ControllerComments{
			Method: "GetIdProveedorXFuncionario",
			Router: `/get_funcionarios_planta`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NaturalezaConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:NominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PorcentajePrimaTecnicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Contratos_x_preliquidacion",
			Router: `/contratos_x_preliquidacion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Contratos_x_preliquidacion_cerrada",
			Router: `/contratos_x_preliquidacion_cerrada`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Resumen",
			Router: `/resumen/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method: "Totales_ss_x_preliquidacion",
			Router: `/totales_x_preliq`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:RelacionParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoConceptoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoNominaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPensionadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:TipoPreliquidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
