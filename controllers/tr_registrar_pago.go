package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// oprations for TrRegistrarPago
type TrRegistrarPago struct {
	beego.Controller
}

func (c *TrRegistrarPago) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRegistrarPago
// @Description create the TrRegistrarPago
// @Param	body		body 	models.TrRegistrarPago	true	"body for TrRegistrarPago content"
// @Success 201 {int} models.TrRegistrarPago
// @Failure 403 body is empty
// @router / [post]
func (c *TrRegistrarPago) Post() {
	var v models.TrRegistrarPago
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRegistrarPago(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = response
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
