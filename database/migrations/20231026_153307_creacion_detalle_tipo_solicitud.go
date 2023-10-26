package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionDetalleTipoSolicitud_20231026_153307 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDetalleTipoSolicitud_20231026_153307{}
	m.Created = "20231026_153307"

	migration.Register("CreacionDetalleTipoSolicitud_20231026_153307", m)
}

// Run the migrations
func (m *CreacionDetalleTipoSolicitud_20231026_153307) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231026_153307_creacion_detalle_tipo_solicitud_up.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionDetalleTipoSolicitud_20231026_153307) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231026_153307_creacion_detalle_tipo_solicitud_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
