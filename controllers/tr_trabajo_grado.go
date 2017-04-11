package controllers

import (
	"github.com/udistrital/Polux_API_Crud/models"
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

// @Title PostTrTrabajoGrado
// @Description create the TrTrabajoGrado
// @Param	body		body 	models.TrabajoGradoTr	true	"body for TrTrabajoGrado content"
// @Success 201 {int} models.TrabajoGradoTr
// @Failure 403 body is empty
// @router / [post]
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
