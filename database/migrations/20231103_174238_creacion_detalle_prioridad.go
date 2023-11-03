package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionDetallePrioridad_20231103_174238 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDetallePrioridad_20231103_174238{}
	m.Created = "20231103_174238"

	migration.Register("CreacionDetallePrioridad_20231103_174238", m)
}

// Run the migrations
func (m *CreacionDetallePrioridad_20231103_174238) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231103_174238_creacion_detalle_prioridad_up.sql")

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
func (m *CreacionDetallePrioridad_20231103_174238) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231103_174238_creacion_detalle_prioridad_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
