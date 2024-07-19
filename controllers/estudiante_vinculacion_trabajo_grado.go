package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Estudiante_Vinculacion_Trabajo_GradoController operations for Estudiante_Vinculacion_Trabajo_Grado
type Estudiante_Vinculacion_Trabajo_GradoController struct {
	beego.Controller
}

// URLMapping ...
func (c *Estudiante_Vinculacion_Trabajo_GradoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Estudiante_Vinculacion_Trabajo_Grado
// @Param	body		body 	models.Estudiante_Vinculacion_Trabajo_Grado	true		"body for Estudiante_Vinculacion_Trabajo_Grado content"
// @Success 201 {object} models.Estudiante_Vinculacion_Trabajo_Grado
// @Failure 403 body is empty
// @router / [post]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Estudiante_Vinculacion_Trabajo_Grado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Estudiante_Vinculacion_Trabajo_Grado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Estudiante_Vinculacion_Trabajo_Grado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Estudiante_Vinculacion_Trabajo_Grado
// @Failure 403
// @router / [get]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) GetAll() {
	//Código del Usuario proveniente del Cliente
	usuario := c.GetString("usuario")

	// Validar parámetros
	if usuario == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Usuario es un parámetro obligatorio"
		c.ServeJSON()
		return
	}

	// Crear el ORM y ejecutar la consulta SQL
	o := orm.NewOrm()
	var result []orm.Params
	query := `SELECT 
			  	vt.id as "Id",
	   		  	vt.usuario as "Usuario",
	   		  	vt.activo as "Activo",
	 	      	vt.fecha_inicio as "FechaInicio",
	          	vt.fecha_fin as "FechaFin",
	          	vt.rol_trabajo_grado as "RolTrabajoGrado",
	          	vt.trabajo_grado as "TrabajoGrado", 
              	tg.titulo as "Titulo", 
              	tg.modalidad as "Modalidad", 
              	tg.estado_trabajo_grado as "EstadoTrabajoGrado", 
              	tg.distincion_trabajo_grado as "DistincionTrabajoGrado", 
              	tg.periodo_academico as "PeriodoAcademico", 
              	tg.objetivo as "Objetivo",
	          	et.estudiante as "Estudiante", 
              	et.estado_estudiante_trabajo_grado as "EstadoEstudiante"
			  FROM academica.vinculacion_trabajo_grado vt
			  LEFT JOIN academica.estudiante_trabajo_grado et
			  ON vt.trabajo_grado = et.trabajo_grado
			  LEFT JOIN academica.trabajo_grado tg
			  ON vt.trabajo_grado = tg.id
			  WHERE vt.activo = true AND vt.usuario = ?`

	_, err := o.Raw(query, usuario).Values(&result)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = "Not found"
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Estudiante_Vinculacion_Trabajo_Grado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Estudiante_Vinculacion_Trabajo_Grado	true		"body for Estudiante_Vinculacion_Trabajo_Grado content"
// @Success 200 {object} models.Estudiante_Vinculacion_Trabajo_Grado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Estudiante_Vinculacion_Trabajo_Grado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Estudiante_Vinculacion_Trabajo_GradoController) Delete() {

}
