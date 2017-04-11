package controllers

import (
	"github.com/udistrital/Polux_API_Crud/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// oprations for TrSolicitud
type TrSolicitudController struct {
	beego.Controller
}

func (c *TrSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrSolicitud
// @Description create the TrSolicitud
// @Param	body		body 	models.TrSolicitud	true	"body for TrSolicitud content"
// @Success 201 {int} models.TrSolicitud
// @Failure 403 body is empty
// @router / [post]
func (c *TrSolicitudController) Post() {
	var v models.TrSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err = models.AddTransaccionSolicitud(&v); err == nil {
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
