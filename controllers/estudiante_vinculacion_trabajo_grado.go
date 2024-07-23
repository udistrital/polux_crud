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
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title Get All
// @Description get Estudiante_Vinculacion_Trabajo_Grado
// @Param	usuario		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EstudianteVinculacionTrabajoGrado
// @Failure 403 :usuario is empty
// @router /:usuario [get]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) GetAll() {
	//Código del Usuario proveniente del Cliente
	usuario := c.Ctx.Input.Param(":usuario")

	// Validar parámetros
	if usuario == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Usuario es un parámetro obligatorio"
		c.ServeJSON()
		return
	}

	v, err := models.GetEstudianteVinculacionTrabajoGrado(usuario)

	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
