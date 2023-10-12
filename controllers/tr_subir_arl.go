package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/polux_crud/models"
)

type TrSubirArlController struct {
	beego.Controller
}

func (c *TrSubirArlController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title PostTrSubirArl
// @Description create the TrSubirArl
// @Param	body		body 	models.TrSubirArl	true	"body for TrSubirArl content"
// @Success 201 {int} models.TrSubirArl
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrSubirArlController) Post() {
	var v models.TrSubirArl
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println(v)
		if response, err := models.AddTransaccionArl(&v); err == nil {
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
