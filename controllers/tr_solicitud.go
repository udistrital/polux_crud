package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/logs"
	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// operations for TrSolicitud
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
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrSolicitudController) Post() {
	var v models.TrSolicitud
	fmt.Println("ENTRANDO A FUNCIÓN")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("ENTRA OBJETO COMPLETO", v)
		if response, err := models.AddTransaccionSolicitud(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": response}
		} else {
			logs.Error(err)
			c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
