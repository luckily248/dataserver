package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"dataserver/models"
)

// get weather information
type AddressResBodyController struct {
	beego.Controller
}

func (c *AddressResBodyController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create AddressResBody
// @Param	body		body 	models.AddressResBody	true		"body for AddressResBody content"
// @Success 200 {int} models.AddressResBody.Id
// @Failure 403 body is empty
// @router / [post]
func (c *AddressResBodyController) Post() {

}

// @Title Get
// @Description get AddressResBody by id
// @Param	id		path 	string	true		"The key for cityid"
// @Success 200 {AddressResBody} models.AddressResBody
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AddressResBodyController) GetOne() {

	cityid:= c.Ctx.Input.Params[":id"]
	fmt.Printf("cityid:%s\n",cityid)
	if cityid != "" {
		ob, err := models.GetOneData(cityid)
		if err != nil {
			c.Data["json"] = err
			fmt.Printf("err:%s\n",err.Error())
		} else {
			c.Data["json"] = ob
			fmt.Printf("ob:%v\n",ob)
		}
	}
	c.ServeJson()
}

// @Title Get All
// @Description get AddressResBody
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AddressResBody
// @Failure 403
// @router / [get]
func (c *AddressResBodyController) GetAll() {

}

// @Title Update
// @Description update the AddressResBody
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AddressResBody	true		"body for AddressResBody content"
// @Success 200 {object} models.AddressResBody
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AddressResBodyController) Put() {

}

// @Title Delete
// @Description delete the AddressResBody
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AddressResBodyController) Delete() {

}
