package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstudianteTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estudiante_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstudianteTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estudiante_trabajo_grado/190"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstudianteTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 192,
		"Estudiante": "20142020001",
		"TrabajoGrado": {
			"Id": 65,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"EstadoEstudianteTrabajoGrado": {
			"Id": 1,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/estudiante_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstudianteTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 192,
		"Estudiante": "20161778011",
		"TrabajoGrado": {
			"Id": 65,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"EstadoEstudianteTrabajoGrado": {
			"Id": 1,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estudiante_trabajo_grado/192", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstudianteTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstudianteTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estudiante_trabajo_grado/192", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstudianteTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
