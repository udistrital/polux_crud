package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// oprations for Formato
type TrFormatoController struct {
	beego.Controller
}

func (c *TrFormatoController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
}

// @Title UpdateTrFormato
// @Description update the TrFormato
// @Param	body		body 	models.TrUpdateFormato	true	"body for TrUpdateFormato content"
// @Success 200 {object} msg
// @Failure 403 :id is not int
// @router / [put]
func (c *TrFormatoController) Put() {
	var v models.TrUpdateFormato
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err = models.UpdateTrFormato(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = "exito"
		} else {
			c.Data["json"] = err.Error()
		}
		c.ServeJSON()
	}
}

// @Title GetFormato
// @Description get Formato by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TrFormato
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TrFormatoController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id", id)
	v, err := models.GetTrFormato(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title PostFormato
// @Description create the Formato
// @Param	body		body 	models.TrSolicitud	true	"body for TrSolicitud content"
// @Success 201 {int} models.TrFormato
// @Failure 403 body is empty
// @router / [post]
func (c *TrFormatoController) Post() {
	var v models.TrFormato
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err = models.AddTrFormato(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
