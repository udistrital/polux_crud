package controllers

import (
	"encoding/json"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// operations for TrRegistrarMateriasPosgrado
type TrRegistrarMateriasPosgrado struct {
	beego.Controller
}

func (c *TrRegistrarMateriasPosgrado) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrRegistrarMateriasPosgrado
// @Description create the TrRegistrarMateriasPosgrado
// @Param	body		body 	models.TrRegistrarMateriasPosgrado	true		"body for TrRegistrarMateriasPosgrado content"
// @Success 201 {int} models.TrRegistrarMateriasPosgrado
// @Failure 403 body is empty
// @router / [post]
func (c *TrRegistrarMateriasPosgrado) Post() {
	var v models.TrRegistrarMateriasPosgrado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if response, err := models.AddTransaccionRegistrarMateriasPosgrado(&v); err == nil {
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
