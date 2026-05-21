package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienDigitalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:BienFisicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:EstadoPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:FavoritoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ImagenPublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionCategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PublicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SubCategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
