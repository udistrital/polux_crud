package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEvaluacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/evaluacion_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEvaluacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEvaluacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEvaluacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEvaluacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/evaluacion_trabajo_grado/80"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEvaluacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEvaluacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEvaluacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEvaluacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 82,
		"Nota": 4.4,
		"VinculacionTrabajoGrado": {
			"Id": 323,
			"Usuario": 0,
			"Activo": false,
			"FechaInicio": "0001-01-01T00:00:00Z",
			"FechaFin": "0001-01-01T00:00:00Z",
			"RolTrabajoGrado": null,
			"TrabajoGrado": null
		},
		"FormatoEvaluacionCarrera": null,
		"Socializacion": null
	}`)

	if response, err := http.Post("http://localhost:9002/v1/evaluacion_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEvaluacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEvaluacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEvaluacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEvaluacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 82,
		"Nota": 4.0,
		"VinculacionTrabajoGrado": {
			"Id": 323,
			"Usuario": 0,
			"Activo": false,
			"FechaInicio": "0001-01-01T00:00:00Z",
			"FechaFin": "0001-01-01T00:00:00Z",
			"RolTrabajoGrado": null,
			"TrabajoGrado": null
		},
		"FormatoEvaluacionCarrera": null,
		"Socializacion": null
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/evaluacion_trabajo_grado/82", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEvaluacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEvaluacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEvaluacionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/evaluacion_trabajo_grado/82", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEvaluacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEvaluacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
