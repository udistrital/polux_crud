package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoAsignaturaTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_asignatura_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoAsignaturaTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_asignatura_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoAsignaturaTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 4,
		"Nombre": "PRUEBA",
		"Descripcion": "Es una prueba",
		"CodigoAbreviacion": "PRUEBA",
		"Activo": true
	}`)

	if response, err := http.Post("http://localhost:9002/v1/estado_asignatura_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoAsignaturaTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 4,
		"Nombre": "PRUEBA",
		"Descripcion": "Es una prueba",
		"CodigoAbreviacion": "PRUEBA",
		"Activo": true
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_asignatura_trabajo_grado/4", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoAsignaturaTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_asignatura_trabajo_grado/4", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
