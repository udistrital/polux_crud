package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrTrabajoGrado
type TrPublicarAsignaturasController struct {
	beego.Controller
}

func (c *TrPublicarAsignaturasController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrPublicarAsignaturas
// @Description create the TrPublicarAsignaturas
// @Param	body		body 	models.TrPublicarAsignaturas	true	"body for TrPublicarAsignaturas content"
// @Success 201 {int} models.TrPublicarAsignaturas
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrPublicarAsignaturasController) Post() {
	var v models.TrPublicarAsignaturas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.AddTransaccionPublicarAsignaturas(&v); err == nil {
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
