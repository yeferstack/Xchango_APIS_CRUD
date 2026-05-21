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

		beego.NSNamespace("/EstadoPublicacion",
			beego.NSInclude(
				&controllers.EstadoPublicacionController{},
			),
		),

		beego.NSNamespace("/Publicacion",
			beego.NSInclude(
				&controllers.PublicacionController{},
			),
		),

		beego.NSNamespace("/ImagenPublicacion",
			beego.NSInclude(
				&controllers.ImagenPublicacionController{},
			),
		),

		beego.NSNamespace("/PublicacionCategoria",
			beego.NSInclude(
				&controllers.PublicacionCategoriaController{},
			),
		),

		beego.NSNamespace("/sub_Categoria",
			beego.NSInclude(
				&controllers.SubCategoriaController{},
			),
		),

		beego.NSNamespace("/BienFisico",
			beego.NSInclude(
				&controllers.BienFisicoController{},
			),
		),

		beego.NSNamespace("/Servicio",
			beego.NSInclude(
				&controllers.ServicioController{},
			),
		),

		beego.NSNamespace("/BienDigital",
			beego.NSInclude(
				&controllers.BienDigitalController{},
			),
		),

		beego.NSNamespace("/Favorito",
			beego.NSInclude(
				&controllers.FavoritoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
