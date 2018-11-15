package controllers

import (
	"encoding/json"

	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrActualizarDocumentoTg
type TrActualizarDocumentoTg struct {
	beego.Controller
}

func (c *TrActualizarDocumentoTg) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrActualizarDocumentoTg
// @Description create the TrActualizarDocumentoTg
// @Param	body		body 	models.TrActualizarDocumentoTg	true		"body for TrActualizarDocumentoTg content"
// @Success 201 {int} models.TrActualizarDocumentoTg
// @Failure 403 body is empty
// @router / [post]
func (c *TrActualizarDocumentoTg) Post() {
	var v models.TrActualizarDocumentoTg
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionActualizarDocumentoTg(&v); err == nil {
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
