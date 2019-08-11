// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/recibo_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_recibo",
			beego.NSInclude(
				&controllers.EstadoReciboController{},
			),
		),

		beego.NSNamespace("/pago_recibo",
			beego.NSInclude(
				&controllers.PagoReciboController{},
			),
		),

		beego.NSNamespace("/tipo_recibo",
			beego.NSInclude(
				&controllers.TipoReciboController{},
			),
		),

		beego.NSNamespace("/recibo",
			beego.NSInclude(
				&controllers.ReciboController{},
			),
		),

		beego.NSNamespace("/tipo_pago",
			beego.NSInclude(
				&controllers.TipoPagoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
