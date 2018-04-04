package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// oprations for TrFormalizarSolicitud
type TrFormalizarSolicitud struct {
	beego.Controller
}

func (c *TrFormalizarSolicitud) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title PostTrFormalizarSolicitud
// @Description create the TrFormalizarSolicitud
// @Param	body		body 	models.TrFormalizarSolicitud	true		"body for TrFormalizarSolicitud content"
// @Success 201 {int} models.TrFormalizarSolicitud
// @Failure 403 body is empty
// @router / [post]
func (c *TrFormalizarSolicitud) Post() {
	var v models.TrFormalizarSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionFormalizarSolicitud(&v); err == nil {
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
