package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoEstudianteTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_estudiante_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoEstudianteTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_estudiante_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoEstudianteTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "PRUEBA",
		"Id": 4,
		"Nombre": "PRUEBA"
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/estado_estudiante_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoEstudianteTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoEstudianteTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoEstudianteTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "PRUEBA",
		"Id": 4,
		"Nombre": "CAMBIO PRUEBA"
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_estudiante_trabajo_grado/4", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoEstudianteTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoEstudianteTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_estudiante_trabajo_grado/4", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoEstudianteTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoEstudianteTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
