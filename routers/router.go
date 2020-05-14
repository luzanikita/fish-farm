// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/nigi4/fish-farm/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/farms",
			beego.NSInclude(
				&controllers.FarmsController{},
			),
			beego.NSRouter("/:id/stats",
				&controllers.FarmsController{}, "get:GetStats",
			),
		),

		beego.NSNamespace("/plans",
			beego.NSInclude(
				&controllers.PlansController{},
			),
		),

		beego.NSNamespace("/conditions",
			beego.NSInclude(
				&controllers.ConditionsController{},
			),
		),

		beego.NSNamespace("/metrics",
			beego.NSInclude(
				&controllers.MetricsController{},
			),
		),

		beego.NSNamespace("/products",
			beego.NSInclude(
				&controllers.ProductsController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/images", "images")

}
