package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	
	//Para concepto_nomina

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:concepto_nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//Para detelle_preliquidacion

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:detalle_preliquidacion"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para estado_concepto_nomina

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_concepto_nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para estado_disponibilidad

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_disponibilidad"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para estado_preliquidacion

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:estado_preliquidacion"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para naturaleza_concepto_nomina

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:naturaleza_concepto_nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para nomina

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para preliquidacion

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:preliquidacion"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para tipo_concepto_nomina

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_concepto_nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para tipo_nomina


	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_nomina"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	//para tipo_preliquidacion

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:tipo_preliquidacion"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
