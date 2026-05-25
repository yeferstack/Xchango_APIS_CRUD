package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CalificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ContactoWhatsAppController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntercambioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:ReporteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
