package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionDetalles_20231026_151404 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDetalles_20231026_151404{}
	m.Created = "20231026_151404"

	migration.Register("CreacionDetalles_20231026_151404", m)
}

// Run the migrations
func (m *CreacionDetalles_20231026_151404) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231026_151404_creacion_detalles_up.sql")

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
func (m *CreacionDetalles_20231026_151404) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231026_151404_creacion_detalles_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
