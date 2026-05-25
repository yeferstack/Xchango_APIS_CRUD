// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers_sistema

import (
	"Xchango_APIS_CRUD/controllers_sistema"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/notificacion",
			beego.NSInclude(
				&controllers_sistema.NotificacionController{},
			),
		),

		beego.NSNamespace("/nivel_acceso",
			beego.NSInclude(
				&controllers_sistema.NivelAccesoController{},
			),
		),

		beego.NSNamespace("/administrador",
			beego.NSInclude(
				&controllers_sistema.AdministradorController{},
			),
		),

		beego.NSNamespace("/administradorpermiso",
			beego.NSInclude(
				&controllers_sistema.AdministradorpermisoController{},
			),
		),

		beego.NSNamespace("/permiso",
			beego.NSInclude(
				&controllers_sistema.PermisoController{},
			),
		),

		beego.NSNamespace("/historialadmin",
			beego.NSInclude(
				&controllers_sistema.HistorialadminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
