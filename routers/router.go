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

		beego.NSNamespace("/docente_cargo",
		beego.NSInclude(
			&controllers.DocenteCargoController{},
			),
		),

		beego.NSNamespace("/funcionario_primatec",
			beego.NSInclude(
				&controllers.FuncionarioPritecController{},
			),
		),

		beego.NSNamespace("/funcionario_proveedor",
		  beego.NSInclude(
		    &controllers.FuncionarioProveedorController{},
		  ),
		),


		beego.NSNamespace("/funcionario_cargo",
					beego.NSInclude(
						&controllers.FuncionarioCargoController{},
					),
				),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/acta_inicio",
			beego.NSInclude(
				&controllers.ActaInicioController{},
			),
		),

		beego.NSNamespace("/argo_ordenadores",
			beego.NSInclude(
				&controllers.ArgoOrdenadoresController{},
			),
		),

		beego.NSNamespace("/relacion_parametro",
			beego.NSInclude(
				&controllers.RelacionParametroController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/estado_civil",
			beego.NSInclude(
				&controllers.EstadoCivilController{},
			),
		),

		beego.NSNamespace("/porcentaje_prima_tecnica",
			beego.NSInclude(
				&controllers.PorcentajePrimaTecnicaController{},
			),
		),

		beego.NSNamespace("/informacion_pensionado",
			beego.NSInclude(
				&controllers.InformacionPensionadoController{},
			),
		),

		beego.NSNamespace("/beneficiario",
			beego.NSInclude(
				&controllers.BeneficiarioController{},
			),
		),

		beego.NSNamespace("/cargo",
			beego.NSInclude(
				&controllers.CargoController{},
			),
		),

		beego.NSNamespace("/categoria_beneficiario",
			beego.NSInclude(
				&controllers.CategoriaBeneficiarioController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/concepto_nomina_por_persona",
			beego.NSInclude(
				&controllers.ConceptoNominaPorPersonaController{},
			),
		),

		beego.NSNamespace("/nomina",
			beego.NSInclude(
				&controllers.NominaController{},
			),
		),

		beego.NSNamespace("/concepto_nomina",
			beego.NSInclude(
				&controllers.ConceptoNominaController{},
			),
		),

		beego.NSNamespace("/pais",
			beego.NSInclude(
				&controllers.PaisController{},
			),
		),

		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),

		beego.NSNamespace("/lugar_ejecucion",
			beego.NSInclude(
				&controllers.LugarEjecucionController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/supervisor_contrato",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),

		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/tipo_pension",
			beego.NSInclude(
				&controllers.TipoPensionController{},
			),
		),

		beego.NSNamespace("/tipo_pensionado",
			beego.NSInclude(
				&controllers.TipoPensionadoController{},
			),
		),

		beego.NSNamespace("/preliquidacion",
			beego.NSInclude(
				&controllers.PreliquidacionController{},
			),
		),

		beego.NSNamespace("/estado_preliquidacion",
			beego.NSInclude(
				&controllers.EstadoPreliquidacionController{},
			),
		),

		beego.NSNamespace("/estado_disponibilidad",
			beego.NSInclude(
				&controllers.EstadoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/estado_concepto_nomina",
			beego.NSInclude(
				&controllers.EstadoConceptoNominaController{},
			),
		),

		beego.NSNamespace("/tipo_nomina",
			beego.NSInclude(
				&controllers.TipoNominaController{},
			),
		),

		beego.NSNamespace("/sustituto",
			beego.NSInclude(
				&controllers.SustitutoController{},
			),
		),

		beego.NSNamespace("/tipo_preliquidacion",
			beego.NSInclude(
				&controllers.TipoPreliquidacionController{},
			),
		),

		beego.NSNamespace("/tipo_concepto_nomina",
			beego.NSInclude(
				&controllers.TipoConceptoNominaController{},
			),
		),

		beego.NSNamespace("/detalle_preliquidacion",
			beego.NSInclude(
				&controllers.DetallePreliquidacionController{},
			),
		),

		beego.NSNamespace("/naturaleza_concepto_nomina",
			beego.NSInclude(
				&controllers.NaturalezaConceptoNominaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
