package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionDetalleTipoSolicitud2_20231103_190542 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDetalleTipoSolicitud2_20231103_190542{}
	m.Created = "20231103_190542"

	migration.Register("CreacionDetalleTipoSolicitud2_20231103_190542", m)
}

// Run the migrations
func (m *CreacionDetalleTipoSolicitud2_20231103_190542) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231103_190542_creacion_detalle_tipo_solicitud_2_up.sql")

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
func (m *CreacionDetalleTipoSolicitud2_20231103_190542) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231103_190542_creacion_detalle_tipo_solicitud_2_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
