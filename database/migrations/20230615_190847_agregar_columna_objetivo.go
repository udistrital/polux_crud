package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarColumnaObjetivo_20230615_190847 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarColumnaObjetivo_20230615_190847{}
	m.Created = "20230615_190847"

	migration.Register("AgregarColumnaObjetivo_20230615_190847", m)
}

// Run the migrations
func (m *AgregarColumnaObjetivo_20230615_190847) Up() {
	file, err := ioutil.ReadFile("../scripts/20230615_190847_agregar_columna_objetivo_up.sql")

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
func (m *AgregarColumnaObjetivo_20230615_190847) Down() {
	file, err := ioutil.ReadFile("../scripts/20230615_190847_agregar_columna_objetivo_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
