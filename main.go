package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/polux_crud/routers"
	"github.com/udistrital/utils_oas/customerror"
	"github.com/udistrital/utils_oas/apiStatusLib"
)

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	_ = orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
}

func main() {
	logPath := "{\"filename\":\""
	logPath += beego.AppConfig.String("logPath")
	logPath += "\"}"
	if err:= logs.SetLogger(logs.AdapterFile, logPath); err != nil{
		if err:= logs.SetLogger("console", ""); err != nil {
			logs.Warn("logPath not set")
		}
	} 
	orm.Debug = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerror.CustomErrorController{})

	apistatus.Init()
	beego.Run()
}
