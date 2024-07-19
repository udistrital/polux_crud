package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/polux_crud/models"
)

// Estudiante_Vinculacion_Trabajo_GradoController operations for Estudiante_Vinculacion_Trabajo_Grado
type Estudiante_Vinculacion_Trabajo_GradoController struct {
	beego.Controller
}

// URLMapping ...
func (c *Estudiante_Vinculacion_Trabajo_GradoController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Estudiante_Vinculacion_Trabajo_Grado by id
// @Param	documento		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EstudianteVinculacionTrabajoGrado
// @Failure 403 :documento is empty
// @router /:documento [get]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) GetOne() {
	documento := c.Ctx.Input.Param(":documento")
	v, err := models.GetEstudianteVinculacionTrabajoGrado(documento)

	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
