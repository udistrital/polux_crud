package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RechazoSustentacion_20230804_004057 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RechazoSustentacion_20230804_004057{}
	m.Created = "20230804_004057"

	migration.Register("RechazoSustentacion_20230804_004057", m)
}

const script = "../scripts/20230804_004057_rechazo_sustentacion"

// Run the migrations
func (m *RechazoSustentacion_20230804_004057) Up() {
	file, err := os.ReadFile(script + "_up.sql")

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
func (m *RechazoSustentacion_20230804_004057) Down() {
	file, err := os.ReadFile(script + "_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
