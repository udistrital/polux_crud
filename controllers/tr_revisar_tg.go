package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRevisarTg
type TrRevisarTg struct {
	beego.Controller
}

func (c *TrRevisarTg) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRevisarTg
// @Description create the TrRevisarTg
// @Param	body		body 	models.TrRevisarTg	true		"body for TrRevisarTg content"
// @Success 201 {int} models.TrRevisarTg
// @Failure 403 body is empty
// @router / [post]
func (c *TrRevisarTg) Post() {
	var v models.TrRevisarTg
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRevisarTg(&v); err == nil {
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
