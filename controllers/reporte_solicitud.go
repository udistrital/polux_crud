package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
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
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": reporteSolicitud}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	c.ServeJSON()
}