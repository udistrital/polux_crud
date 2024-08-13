package controllers

import (
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
		c.Data["json"] = reporteGeneral
		c.Ctx.Output.SetStatus(201)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": err.Error()}
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()
}
