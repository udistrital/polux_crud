package controllers

import (
	"github.com/udistrital/Polux_API_Crud/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

// oprations for Formato
type TrFormatoController struct {
	beego.Controller
}

func (c *TrFormatoController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

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
