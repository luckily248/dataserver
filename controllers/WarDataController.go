package controllers

import (
	"dataserver/models"
	"fmt"

	"github.com/astaxie/beego"
)

// get weather information
type WarDataController struct {
	beego.Controller
}

func (c *WarDataController) URLMapping() {
	c.Mapping("NewWar", c.NewWar)
	c.Mapping("GetWar", c.GetWar)
	c.Mapping("Bot", c.Bot)
}

// @Title Post
// @Description post new WarDataModel by TeamA TeamB
// @Param	TeamA		form 	string	true		"new war TeamA"
// @Param	TeamB		form 	string	true		"new war TeamB"
// @Success 200 {string} string
// @Failure 404  not found
// @router /newwar [post]
func (c *WarDataController) NewWar() {
	wardata := models.WarDataModel{}
	if err := c.ParseForm(&wardata); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = err.Error()
		fmt.Printf("err:%s\n", err.Error())
	}
	fmt.Printf("teama:%s\n", wardata.TeamA)
	fmt.Printf("teamb:%s\n", wardata.TeamB)
	if wardata.TeamA == "" || wardata.TeamB == "" {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = string("someteam empty")
	} else {
		id, err := models.AddWarData(wardata.TeamA, wardata.TeamB)
		if err != nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
			fmt.Printf("err:%s\n", err.Error())
		} else {
			c.Data["json"] = id
			fmt.Printf("id:%s\n", id)
		}
	}
	c.ServeJson()
}

// @Title Get
// @Description get WarDataModel by clanname
// @Param	TeamA		form 	string	true		"new war TeamA"
// @Param	TeamB		form 	string	true		"new war TeamB"
// @Success 200 {object} models.WarDataModel
// @Failure 404  not found
// @Failure 403 :clanname is empty
// @router /:clanname [get]
func (c *WarDataController) GetWar() {

	clanname := c.Ctx.Input.Params[":clanname"]
	fmt.Printf("clanname:%s\n", clanname)
	if clanname == "" {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = string("is empty")
	} else {
		ob, err := models.GetWarData(clanname)
		if err != nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
			fmt.Printf("err:%s\n", err.Error())
		} else {
			c.Data["json"] = ob
			fmt.Printf("ob:%v\n", ob)
		}
	}
	c.ServeJson()
}

// @Title Post
// @Description get rep from callerbot
// @Param	clanname		path 	string	true		"The key for clan"
// @Success 200 {object} models.WarDataModel
// @Failure 404  not found
// @router /bot [post]
func (c *WarDataController) Bot() {

	clanname := c.Ctx.Input.Params[":clanname"]
	fmt.Printf("clanname:%s\n", clanname)
	if clanname == "" {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = string("is empty")
	} else {
		ob, err := models.GetWarData(clanname)
		if err != nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
			fmt.Printf("err:%s\n", err.Error())
		} else {
			c.Data["json"] = ob
			fmt.Printf("ob:%v\n", ob)
		}
	}
	c.ServeJson()
}
