package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoNominaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ContratoPreliquidacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DetallePreliquidacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
