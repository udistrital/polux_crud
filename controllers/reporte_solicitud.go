package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/polux_crud/models"
)

// Reporte_solicitudController operations for Reporte_solicitud
type ReporteSolicitudController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReporteSolicitudController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get ReporteSolicitud
// @Success 201 {object} models.ReporteSolicitud
// @Failure 403
// @router / [get]
func (c *ReporteSolicitudController) GetAll() {
	reporteSolicitud, err := models.GetReporteSolicitud()
	if err == nil {
		c.Data["json"] = reporteSolicitud
		c.Ctx.Output.SetStatus(201)
	} else {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": err.Error()}
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()
}
