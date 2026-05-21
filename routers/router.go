// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
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

		beego.NSNamespace("/notificacion",
			beego.NSInclude(
				&controllers.NotificacionController{},
			),
		),

		beego.NSNamespace("/nivel_acceso",
			beego.NSInclude(
				&controllers.NivelAccesoController{},
			),
		),

		beego.NSNamespace("/administrador",
			beego.NSInclude(
				&controllers.AdministradorController{},
			),
		),

		beego.NSNamespace("/administradorpermiso",
			beego.NSInclude(
				&controllers.AdministradorpermisoController{},
			),
		),

		beego.NSNamespace("/permiso",
			beego.NSInclude(
				&controllers.PermisoController{},
			),
		),

		beego.NSNamespace("/historialadmin",
			beego.NSInclude(
				&controllers.HistorialadminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
