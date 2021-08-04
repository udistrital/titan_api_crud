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

		beego.NSNamespace("/nomina",
			beego.NSInclude(
				&controllers.NominaController{},
			),
		),

		beego.NSNamespace("/preliquidacion",
			beego.NSInclude(
				&controllers.PreliquidacionController{},
			),
		),

		beego.NSNamespace("/naturaleza_concepto_nomina",
			beego.NSInclude(
				&controllers.NaturalezaConceptoNominaController{},
			),
		),

		beego.NSNamespace("/concepto_nomina",
			beego.NSInclude(
				&controllers.ConceptoNominaController{},
			),
		),

		beego.NSNamespace("/tipo_concepto_nomina",
			beego.NSInclude(
				&controllers.TipoConceptoNominaController{},
			),
		),

		beego.NSNamespace("/estado_concepto_nomina",
			beego.NSInclude(
				&controllers.EstadoConceptoNominaController{},
			),
		),

		beego.NSNamespace("/detalle_preliquidacion",
			beego.NSInclude(
				&controllers.DetallePreliquidacionController{},
			),
		),

		beego.NSNamespace("/estado_disponibilidad",
			beego.NSInclude(
				&controllers.EstadoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/tipo_preliquidacion",
			beego.NSInclude(
				&controllers.TipoPreliquidacionController{},
			),
		),

		beego.NSNamespace("/tipo_nomina",
			beego.NSInclude(
				&controllers.TipoNominaController{},
			),
		),

		beego.NSNamespace("/estado_preliquidacion",
			beego.NSInclude(
				&controllers.EstadoPreliquidacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
