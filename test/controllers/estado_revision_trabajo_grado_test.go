package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoRevisionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_revision_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoRevisionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_revision_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoRevisionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 4,
		"Nombre": "PRUEBA"
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/estado_revision_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoRevisionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 4,
		"Nombre": "CAMBIO PRUEBA"
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_revision_trabajo_grado/4", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoRevisionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoRevisionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_revision_trabajo_grado/4", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoRevisionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
