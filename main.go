package main

import (
	_ "dataserver/docs"
	_ "dataserver/routers"
	"github.com/astaxie/beego"
	"dataserver/collector"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	collector.Run()
	beego.Run()
	
}
