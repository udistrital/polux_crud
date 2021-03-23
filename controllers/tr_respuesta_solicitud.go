package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrTrabajoGrado
type TrRespuestaSolicitudController struct {
	beego.Controller
}

func (c *TrRespuestaSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRespuestaSolicitud
// @Description create the TrRespuestaSolicitud
// @Param	body		body 	models.TrRespuestaSolicitud	true	"body for TrRespuestaSolicitud content"
// @Success 201 {int} models.TrRespuestaSolicitud
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrRespuestaSolicitudController) Post() {
	var v models.TrRespuestaSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRespuestaSolicitud(&v); err == nil {
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
