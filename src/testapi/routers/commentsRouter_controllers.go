package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "AddStudent",
            Router: "/AddStudent/:ClassId/:StudentId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:ClassController"] = append(beego.GlobalControllerRouter["testapi/controllers:ClassController"],
        beego.ControllerComments{
            Method: "GetAllStudentOfClass",
            Router: "/GetAllStudentOfClass/:ClassId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetAllClassesOfStudent",
            Router: "/GetAllClassesOfStudent/:StudentId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["testapi/controllers:StudentController"] = append(beego.GlobalControllerRouter["testapi/controllers:StudentController"],
        beego.ControllerComments{
            Method: "AddClass",
            Router: "/addClass/:StudentId/:ClassId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
