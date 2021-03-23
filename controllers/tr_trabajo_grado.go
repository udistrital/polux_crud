package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrTrabajoGrado
type TrTrabajoGradoController struct {
	beego.Controller
}

func (c *TrTrabajoGradoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrTrabajoGrado
// @Description create the TrTrabajoGrado
// @Param	body		body 	models.TrTrabajoGrado	true	"body for TrTrabajoGrado content"
// @Success 201 {int} models.TrTrabajoGrado
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrTrabajoGradoController) Post() {
	var v models.TrTrabajoGrado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.AddTransaccionTrabajoGrado(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
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
