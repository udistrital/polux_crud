package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/polux_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"
)

/*func init() {
	if errorDriver := orm.RegisterDriver("postgres", orm.DRPostgres); errorDriver == nil {
		if errorDB := orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=require&search_path="+beego.AppConfig.String("PGschemas")+""); errorDB == nil {
			beego.Info("Success")
		} else {
			beego.Info(errorDB)
		}
	} else {
		beego.Info(errorDriver)
	}
}*/

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		beego.AppConfig.String("PGpass")+"@"+
		beego.AppConfig.String("PGurls")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()
	beego.Run()
}
