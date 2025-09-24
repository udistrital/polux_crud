package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllFormatoEvaluacionCarrera(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/formato_evaluacion_carrera/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllFormatoEvaluacionCarrera Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllFormatoEvaluacionCarrera Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllFormatoEvaluacionCarrera:", err.Error())
		t.Fail()
	}
}

func TestPostFormatoEvaluacionCarrera(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoProyecto": 2,
		"FechaFin": "2020-12-21T14:03:14Z",
		"FechaInicio": "2020-12-24T14:03:14Z",
		"Formato": {
		  "FechaRealizacion": "2020-12-21T14:03:14Z",
		  "Id": 1,
		  "Introduccion": "INTRO",
		  "Nombre": "PRUEBA"
		},
		"Id": 1,
		"Modalidad": {
		  "Activa": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 2,
		  "Nombre": "string"
		}
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/formato_evaluacion_carrera/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostFormatoEvaluacionCarrera Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostFormatoEvaluacionCarrera Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostFormatoEvaluacionCarrera:", err.Error())
		t.Fail()
	}
}

func TestGetOneFormatoEvaluacionCarrera(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/formato_evaluacion_carrera/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneFormatoEvaluacionCarrera Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneFormatoEvaluacionCarrera Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneFormatoEvaluacionCarrera:", err.Error())
		t.Fail()
	}
}

func TestPutFormatoEvaluacionCarrera(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoProyecto": 2,
		"FechaFin": "2020-12-21T14:03:14Z",
		"FechaInicio": "2020-12-24T14:03:14Z",
		"Formato": {
		  "FechaRealizacion": "2020-12-21T14:03:14Z",
		  "Id": 1,
		  "Introduccion": "INTRO",
		  "Nombre": "CAMBIO PRUEBA"
		},
		"Id": 1,
		"Modalidad": {
		  "Activa": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 2,
		  "Nombre": "string"
		}
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/formato_evaluacion_carrera/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutFormatoEvaluacionCarrera Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutFormatoEvaluacionCarrera Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteFormatoEvaluacionCarrera(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/formato_evaluacion_carrera/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteFormatoEvaluacionCarrera Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteFormatoEvaluacionCarrera Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
