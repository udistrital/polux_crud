package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDistincionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/distincion_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDistincionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDistincionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDistincionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneDistincionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/distincion_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDistincionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDistincionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDistincionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostDistincionTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 3,
        "Nombre": "PRUEBAS",
        "Descripcion": "Es una prueba",
        "CodigoAbreviacion": "PRUEBA",
        "Activo": false
    }`)

	if response, err := http.Post("http://localhost:9002/v1/distincion_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDistincionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDistincionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDistincionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutDistincionTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 3,
        "Nombre": "PRUEBAS",
        "Descripcion": "Es una prueba",
        "CodigoAbreviacion": "PRUEBA",
        "Activo": false
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/distincion_trabajo_grado/3", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDistincionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDistincionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDistincionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/distincion_trabajo_grado/3", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDistincionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDistincionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
