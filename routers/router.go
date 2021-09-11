// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/titan_api_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/detalle_preliquidacion",
			beego.NSInclude(
				&controllers.DetallePreliquidacionController{},
			),
		),

		beego.NSNamespace("/preliquidacion",
			beego.NSInclude(
				&controllers.PreliquidacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
