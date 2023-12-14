package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DocumentoCorreccion_20231214_004230 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DocumentoCorreccion_20231214_004230{}
	m.Created = "20231214_004230"

	migration.Register("DocumentoCorreccion_20231214_004230", m)
}

// Run the migrations
func (m *DocumentoCorreccion_20231214_004230) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20231214_004230_documento_correccion_up.sql")

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
func (m *DocumentoCorreccion_20231214_004230) Down() {
	file, err := os.ReadFile("../scripts/20231214_004230_documento_correccion_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
