package main

import (
	_ "github.com/nigi4/fish-farm/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	beego.Run()
}

