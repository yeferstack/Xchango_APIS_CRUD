package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:AdministradorpermisoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialadminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NivelAccesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PermisoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
