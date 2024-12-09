package controllers

import (
	"encoding/json"

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
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description get ReporteSolicitud
// @Param	body		body 	models.FiltrosReporte	true		"body for FiltrosReporte content"
// @Success 201 {object} models.ReporteSolicitud
// @Failure 403
// @router / [post]
func (c *ReporteSolicitudController) Post() {
	var v models.FiltrosReporte
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		reporteSolicitud, err := models.GetReporteSolicitud(&v)
		if err == nil {
			c.Data["json"] = reporteSolicitud
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
