package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 24,
		"Nombre": "PRUEBA"
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/estado_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 24,
		"Nombre": "CAMBIO PRUEBA"
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_trabajo_grado/24", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_trabajo_grado/24", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
