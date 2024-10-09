package controllers

import (
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
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get ReporteGeneral
// @Success 201 {object} models.ReporteGeneral
// @Failure 403
// @router / [get]
func (c *ReporteGeneralController) GetAll() {
	reporteGeneral, err := models.GetReporteGeneral()
	if err == nil {
		if reporteGeneral["Success"].(bool) {
			c.Data["json"] = reporteGeneral
			c.Ctx.Output.SetStatus(reporteGeneral["Status"].(int))
		} else {
			c.Data["json"] = reporteGeneral
			c.Ctx.Output.SetStatus(reporteGeneral["Status"].(int))
		}
	} else {
		logs.Error(err)
		c.Data["message"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	c.ServeJSON()
}
