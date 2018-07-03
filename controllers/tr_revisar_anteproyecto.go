package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRevisarAnteproyecto
type TrRevisarAnteproyecto struct {
	beego.Controller
}

func (c *TrRevisarAnteproyecto) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRevisarAnteproyecto
// @Description create the TrRevisarAnteproyecto
// @Param	body		body 	models.TrRevisarAnteproyecto	true		"body for TrRevisarAnteproyecto content"
// @Success 201 {int} models.TrRevisarAnteproyecto
// @Failure 403 body is empty
// @router / [post]
func (c *TrRevisarAnteproyecto) Post() {
	var v models.TrRevisarAnteproyecto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRevisarAnteproyecto(&v); err == nil {
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
