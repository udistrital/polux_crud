package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrSeleccionAdmitidos
type TrSeleccionAdmitidosController struct {
	beego.Controller
}

func (c *TrSeleccionAdmitidosController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrSeleccionAdmitidos
// @Description create the TrSeleccionAdmitidos
// @Param	body		body 	models.TrSeleccionAdmitidos	true	"body for TrSeleccionAdmitidos content"
// @Success 201 {int} models.TrSeleccionAdmitidos
// @Failure 403 body is empty
// @router / [post]
func (c *TrSeleccionAdmitidosController) Post() {
	var v models.TrSeleccionAdmitidos
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.TransaccionSeleccionAdmitidos(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = alerta
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
