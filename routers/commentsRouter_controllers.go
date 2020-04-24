package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ConditionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:FarmsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:MetricsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:PlansController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:RequirementsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/nigi4/fish-farm/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
