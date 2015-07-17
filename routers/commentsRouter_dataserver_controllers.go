package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:UserController"] = append(beego.GlobalControllerRouter["dataserver/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

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

	beego.GlobalControllerRouter["dataserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dataserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dataserver/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

}
