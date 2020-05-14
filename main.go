package main

import (
	_ "github.com/nigi4/fish-farm/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
    "github.com/astaxie/beego/session"

	_ "github.com/lib/pq"
)

var GlobalSessions *session.Manager

func main() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.BConfig.WebConfig.Session.SessionOn = true
    sessionConfig := &session.ManagerConfig{
		CookieName: "gosessionid",
		Gclifetime: 3600,
	}
	GlobalSessions, _ := session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()

	beego.Run()
}
