package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["organization2/controllers:UnitController"] = append(beego.GlobalControllerRouter["organization2/controllers:UnitController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UnitController"] = append(beego.GlobalControllerRouter["organization2/controllers:UnitController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UnitController"] = append(beego.GlobalControllerRouter["organization2/controllers:UnitController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UnitController"] = append(beego.GlobalControllerRouter["organization2/controllers:UnitController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserSortedByPage",
            Router: "/get-by-page/:page/:size/:type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["organization2/controllers:UserController"] = append(beego.GlobalControllerRouter["organization2/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUnitOfUser",
            Router: "/get-unit/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
