package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDetalleTrabajoGrado_20230712_011240 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDetalleTrabajoGrado_20230712_011240{}
	m.Created = "20230712_011240"

	migration.Register("CrearTablaDetalleTrabajoGrado_20230712_011240", m)
}

// Run the migrations
func (m *CrearTablaDetalleTrabajoGrado_20230712_011240) Up() {
	file, err := ioutil.ReadFile("../scripts/20230712_011240_crear_tabla_detalle_trabajo_grado_up.sql")

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
func (m *CrearTablaDetalleTrabajoGrado_20230712_011240) Down() {
	file, err := ioutil.ReadFile("../scripts/20230712_011240_crear_tabla_detalle_trabajo_grado_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
