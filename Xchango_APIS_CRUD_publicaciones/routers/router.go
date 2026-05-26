// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	 "XCHANGO_APIS_CRUD/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/publicacion",
			beego.NSInclude(
				&controllers.PublicacionController{},
			),
		),

		beego.NSNamespace("/imagenpublicacion",
			beego.NSInclude(
				&controllers.ImagenpublicacionController{},
			),
		),

		beego.NSNamespace("/publicacioncategoria",
			beego.NSInclude(
				&controllers.PublicacioncategoriaController{},
			),
		),

		beego.NSNamespace("/sub_categoria",
			beego.NSInclude(
				&controllers.SubCategoriaController{},
			),
		),

		beego.NSNamespace("/bienfisico",
			beego.NSInclude(
				&controllers.BienfisicoController{},
			),
		),

		beego.NSNamespace("/servicio",
			beego.NSInclude(
				&controllers.ServicioController{},
			),
		),

		beego.NSNamespace("/biendigital",
			beego.NSInclude(
				&controllers.BiendigitalController{},
			),
		),

		beego.NSNamespace("/favorito",
			beego.NSInclude(
				&controllers.FavoritoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
