package controllers

import (
	"github.com/udistrital/Polux_API_Crud/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// oprations for TipoVinculacion
type TrSolicitudController struct {
	beego.Controller
}

func (c *TrSolicitudController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *TrSolicitudController) Post() {
	var v models.TrSolicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//fmt.Println(*v.Tg)
		//	var t models.TrabajoGrado
		//	t = *v.Tg
		if _, err = models.AddTransaccionSolicitud(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
			/*for i := 0; i < len(*v.MateriasSolicitadas); i++ {
				var asignatura models.AsignaturaInscrita
				asignatura = &v.MateriasSolicitadas[i]
				if _, err = models.AddAsignaturaInscrita(&asignatura); err == nil {
					c.Ctx.Output.SetStatus(201)
					c.Data["json"] = v
				} else {
					c.Data["json"] = err.Error()
				}
			}*/

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
