package routers

import (
	"github.com/nigi4/fish-farm/fish-farm-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
