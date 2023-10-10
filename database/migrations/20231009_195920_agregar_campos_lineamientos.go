package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposLineamientos_20231009_195920 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposLineamientos_20231009_195920{}
	m.Created = "20231009_195920"

	migration.Register("AgregarCamposLineamientos_20231009_195920", m)
}

// Run the migrations
func (m *AgregarCamposLineamientos_20231009_195920) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231009_195920_agregar_campos_lineamientos_up.sql")

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
func (m *AgregarCamposLineamientos_20231009_195920) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231009_195920_agregar_campos_lineamientos_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
