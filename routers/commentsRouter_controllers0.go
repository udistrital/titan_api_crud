package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"],
		beego.ControllerComments{
			Method:           "ConsultarIDProveedor",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method:           "Resumen",
			Router:           `/resumen`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method:           "Resumen",
			Router:           `/resumen`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoPorPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoPorPersonaController"],
		beego.ControllerComments{
			Method:           "NovedadesActivas",
			Router:           `novedades_activas/:id`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
		beego.ControllerComments{
			Method:           "ConsultarAsignacionBasica",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
		beego.ControllerComments{
			Method:           "ConsultarCedulaDocente",
			Router:           `/consultarCedulaDocente`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"],
		beego.ControllerComments{
			Method:           "ConsultarAsignacionBasica",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

		beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"],
					beego.ControllerComments{
					Method:           "ConsultarPrimaTecnica",
					Router:           `/`,
					AllowHTTPMethods: []string{"post"},
					Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
		beego.ControllerComments{
			Method:           "Pensionado_datos",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
}
