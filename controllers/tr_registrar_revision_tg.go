package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRegistrarRevisionTg
type TrRegistrarRevisionTgController struct {
	beego.Controller
}

func (c *TrRegistrarRevisionTgController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRegistrarRevisionTg
// @Description create the TrRegistrarRevisionTg
// @Param	body		body 	models.TrRegistrarRevisionTg	true		"body for TrRegistrarRevisionTg content"
// @Success 201 {int} models.TrRegistrarRevisionTg
// @Failure 403 body is empty
// @router / [post]
func (c *TrRegistrarRevisionTgController) Post() {
	var v models.TrRegistrarRevisionTg
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRegistrarRevisionTg(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = response
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
		beego.Error(err)
		c.Abort("400")
	}
	c.ServeJSON()
}
