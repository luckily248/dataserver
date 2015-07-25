package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dataserver/controllers:CityContentController"] = append(beego.GlobalControllerRouter["dataserver/controllers:CityContentController"],
		beego.ControllerComments{
			"GetOneCityContent",
			`/citycontent`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:CityContentController"] = append(beego.GlobalControllerRouter["dataserver/controllers:CityContentController"],
		beego.ControllerComments{
			"GetOneCity",
			`/:name`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ShowWeatherController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

}
