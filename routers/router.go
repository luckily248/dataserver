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
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/addressresbody",
			beego.NSInclude(
				&controllers.AddressResBodyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
