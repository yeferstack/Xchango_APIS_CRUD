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

		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/Credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/Crear_contrasena",
			beego.NSInclude(
				&controllers.CrearContrasenaController{},
			),
		),

		beego.NSNamespace("/HistorialContrasena",
			beego.NSInclude(
				&controllers.HistorialContrasenaController{},
			),
		),

		beego.NSNamespace("/RecuperacionContrasena",
			beego.NSInclude(
				&controllers.RecuperacionContrasenaController{},
			),
		),

		beego.NSNamespace("/Sesion",
			beego.NSInclude(
				&controllers.SesionController{},
			),
		),

		beego.NSNamespace("/IntentoLogin",
			beego.NSInclude(
				&controllers.IntentoLoginController{},
			),
		),

		beego.NSNamespace("/Perfil",
			beego.NSInclude(
				&controllers.PerfilController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
