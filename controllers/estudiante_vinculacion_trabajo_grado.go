package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
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
// @Failure 404 not found resource
// @router / [get]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) GetAll() {
	// Obtener los parámetros enviados desde el cliente
	usuario := c.GetString("usuario")
	proyectos := c.GetString("proyectos")

	if usuario != "" && proyectos == "" {
		// Caso 1: Se envía solo el parámetro `usuario`
		v, err := models.GetEstudianteVinculacionTrabajoGrado(usuario)

		if err != nil {
			logs.Error(err)
			c.Data["message"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		} else {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
		}
	} else if proyectos != "" && usuario == "" {
		// Caso 2: Se envía solo el parámetro `proyectos`
		v, err := models.GetProyectoVinculacionTrabajoGrado(proyectos)

		if err != nil {
			logs.Error(err)
			c.Data["message"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		} else {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
		}
	} else {
		// Caso 3: Parámetros inválidos o ninguno presente
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Debe enviar al menos el campo de usuario o proyectos"
		c.ServeJSON()
		return
	}

	// Devolver la respuesta al cliente en JSON
	c.ServeJSON()
}
