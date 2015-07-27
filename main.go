package main

import (
	_ "dataserver/docs"
	_ "dataserver/routers"
	"github.com/astaxie/beego"
	"dataserver/collector"
	"github.com/astaxie/beego/toolbox"
	"dataserver/healthcheck"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	toolbox.AddHealthCheck("database",&healthcheck.DatabaseCheck{})
	toolbox.AddHealthCheck("citycontent",&healthcheck.CitycontentCheck{})
	collector.Run()
	beego.Run()
}
