// @APIVersion 1.0.0
// @Title DataServer API
// @Description a dataserver api for data
// @Contact luck248@163.com
// @TermsOfServiceUrl http://luckily.cc/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dataserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1", //v1 :版本号
		beego.NSNamespace("/weather", //天气API  e.g. http://127.0.0.1:8888/v1/weather/:id
			beego.NSInclude(
				&controllers.ShowWeatherController{},
			),
		),
		beego.NSNamespace("/city", //城市id API  e.g. http://127.0.0.1:8888/v1/city/:name
			beego.NSInclude(
				&controllers.CityContentController{},
			),
		),
		beego.NSNamespace("/cocserver", //cocserver API  e.g. http://127.0.0.1:8888/v1/cocserver/:warid
			beego.NSInclude(
				&controllers.WarDataController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
