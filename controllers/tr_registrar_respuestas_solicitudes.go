package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrSeleccionAdmitidos
type TrRegistrarRespuestasSolicitudesController struct {
	beego.Controller
}

func (c *TrRegistrarRespuestasSolicitudesController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrSeleccionAdmitidos
// @Description create the TrSeleccionAdmitidos
// @Param	body		body 	models.TrSeleccionAdmitidos	true	"body for TrSeleccionAdmitidos content"
// @Success 201 {int} models.TrSeleccionAdmitidos
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrRegistrarRespuestasSolicitudesController) Post() {
	var v models.TrRegistrarRespuestasSolicitudes
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.TransaccionRegistrarRespuestasSolicitudes(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = response
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
