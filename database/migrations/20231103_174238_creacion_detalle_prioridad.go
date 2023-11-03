package main

import (
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

}

// Reverse the migrations
func (m *CreacionDetallePrioridad_20231103_174238) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
