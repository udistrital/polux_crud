package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRegistrarNota
type TrRegistrarNota struct {
	beego.Controller
}

func (c *TrRegistrarNota) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRegistrarNota
// @Description create the TrRegistrarNota
// @Param	body		body 	models.TrRegistrarNota	true		"body for TrRegistrarNota content"
// @Success 201 {int} models.TrRegistrarNota
// @Failure 403 body is empty
// @router / [post]
func (c *TrRegistrarNota) Post() {
	var v models.TrRegistrarNota
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRegistrarNota(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = response
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
