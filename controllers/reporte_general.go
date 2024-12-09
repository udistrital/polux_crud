package controllers

import (
	"encoding/json"

	"github.com/beego/beego/logs"
	"github.com/udistrital/polux_crud/models"

	"github.com/astaxie/beego"
)

// ReporteGeneralController operations for reporte_general
type ReporteGeneralController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReporteGeneralController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description get ReporteGeneral
// @Param	body		body 	models.FiltrosReporte	true		"body for FiltrosReporte content"
// @Success 201 {object} models.ReporteGeneral
// @Failure 400
// @router / [post]
func (c *ReporteGeneralController) Post() {
	var v models.FiltrosReporte
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		reporteGeneral, err := models.GetReporteGeneral(&v)
		if err == nil {
			c.Data["json"] = reporteGeneral
			c.Ctx.Output.SetStatus(201)
		} else {
			logs.Error(err)
			c.Data["message"] = "Error service Post: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
