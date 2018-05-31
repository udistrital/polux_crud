package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRegistrarActaSeguimiento
type TrRegistrarActaSeguimientoController struct {
	beego.Controller
}

func (c *TrRegistrarActaSeguimientoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRegistrarActaSeguimiento
// @Description create the TrRegistrarActaSeguimiento para registrar actas de seguimiento en la modalidad de pasantia
// @Param	body		body 	models.TrRegistrarActaSeguimiento	true	"body for TrRegistrarActaSeguimiento content"
// @Success 201 {int} models.TrRegistrarActaSeguimiento
// @Failure 403 body is empty
// @router / [post]
func (c *TrRegistrarActaSeguimientoController) Post() {
	var v models.TrRegistrarActaSeguimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.RegistrarActaSeguimiento(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}