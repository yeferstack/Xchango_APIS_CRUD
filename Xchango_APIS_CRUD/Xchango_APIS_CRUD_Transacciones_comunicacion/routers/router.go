// @APIVersion 1.0.0
// @Title Xchango Transacciones API
// @Description API para el esquema Transacciones_comunicacion
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Xchango_APIS_CRUD/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Intercambio",
			beego.NSInclude(
				&controllers.IntercambioController{},
			),
		),

		beego.NSNamespace("/ContactoWhatsApp",
			beego.NSInclude(
				&controllers.ContactoWhatsAppController{},
			),
		),

		beego.NSNamespace("/Calificacion",
			beego.NSInclude(
				&controllers.CalificacionController{},
			),
		),

		beego.NSNamespace("/Reporte",
			beego.NSInclude(
				&controllers.ReporteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
