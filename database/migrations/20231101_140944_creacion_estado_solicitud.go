package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionEstadoSolicitud_20231101_140944 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionEstadoSolicitud_20231101_140944{}
	m.Created = "20231101_140944"

	migration.Register("CreacionEstadoSolicitud_20231101_140944", m)
}

// Run the migrations
func (m *CreacionEstadoSolicitud_20231101_140944) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231101_140944_creacion_estado_solicitud_up.sql")

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
func (m *CreacionEstadoSolicitud_20231101_140944) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231101_140944_creacion_estado_solicitud_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
