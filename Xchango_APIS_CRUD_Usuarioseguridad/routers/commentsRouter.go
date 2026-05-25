package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CrearContrasenaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:HistorialContrasenaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:IntentoLoginController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:RecuperacionContrasenaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:SesionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["Xchango_APIS_CRUD/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
