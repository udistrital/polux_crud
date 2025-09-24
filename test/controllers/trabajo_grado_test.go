package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/trabajo_grado/197"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 209,
		"Titulo": "PRUEBA",
		"Modalidad": {
			"Id": 4,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activa": false
		},
		"EstadoTrabajoGrado": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"DistincionTrabajoGrado": null,
		"PeriodoAcademico": "2023-3",
		"Objetivo": "Es una prueba"
	}`)

	if response, err := http.Post("http://localhost:9002/v1/trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 209,
		"Titulo": "PRUEBA",
		"Modalidad": {
			"Id": 4,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activa": false
		},
		"EstadoTrabajoGrado": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"DistincionTrabajoGrado": null,
		"PeriodoAcademico": "2023-3",
		"Objetivo": "Es una prueba"
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/trabajo_grado/209", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/trabajo_grado/209", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
