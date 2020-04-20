package routers

import (
	"github.com/nigi4/fish-farm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
