package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"] = append(beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"] = append(beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"] = append(beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"] = append(beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"] = append(beego.GlobalControllerRouter["dataserver/controllers:AddressResBodyController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
