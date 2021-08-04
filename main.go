package main

import (
	_ "titan_api_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/customerrorv2"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
		  "content-type",
		  "accept",
		  "origin",
		  "authorization",
		  "x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	  }))
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	beego.Run()
}

