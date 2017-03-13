package controllers

import (
	"Polux_API_Crud/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// oprations for TgTr
type TrTrabajoGradoController struct {
	beego.Controller
}

func (c *TrTrabajoGradoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *TrTrabajoGradoController) Post() {
	var v models.TrabajoGradoTr

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if _, err := models.AddTrabajoGradoTr(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
