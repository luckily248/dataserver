package main

import (
	"./collector"
	_ "./docs"
	"./healthcheck"
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	toolbox.AddHealthCheck("database", &healthcheck.DatabaseCheck{})
	toolbox.AddHealthCheck("citycontent", &healthcheck.CitycontentCheck{})
	collector.Run()
	beego.Run()
}
