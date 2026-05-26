package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BiendigitalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:BienfisicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ImagenpublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:PublicacioncategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["XCHANGO_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
