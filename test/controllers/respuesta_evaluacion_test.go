package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRespuestaEvaluacion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_evaluacion/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRespuestaEvaluacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRespuestaEvaluacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRespuestaEvaluacion:", err.Error())
		t.Fail()
	}
}

func TestPostRespuestaEvaluacion(t *testing.T) {
	body := []byte(`{
		"Id": 1,
		"Justificacion": "Prueba",
		"RespuestaFormato": {
			"Id": 1,
			"Orden": 0,
			"Valoracion": 0,
			"Respuesta": null,
			"PreguntaFormato": null
		},
		"EvaluacionTrabajoGrado": {
			"Id": 81,
			"Nota": 0,
			"VinculacionTrabajoGrado": null,
			"FormatoEvaluacionCarrera": null,
			"Socializacion": null
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/respuesta_evaluacion/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRespuestaEvaluacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRespuestaEvaluacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRespuestaEvaluacion:", err.Error())
		t.Fail()
	}
}

func TestGetOneRespuestaEvaluacion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_evaluacion/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRespuestaEvaluacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRespuestaEvaluacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRespuestaEvaluacion:", err.Error())
		t.Fail()
	}
}

func TestPutRespuestaEvaluacion(t *testing.T) {
	body := []byte(`{
		"Id": 1,
		"Justificacion": "Prueba",
		"RespuestaFormato": {
			"Id": 1,
			"Orden": 0,
			"Valoracion": 0,
			"Respuesta": null,
			"PreguntaFormato": null
		},
		"EvaluacionTrabajoGrado": {
			"Id": 81,
			"Nota": 0,
			"VinculacionTrabajoGrado": null,
			"FormatoEvaluacionCarrera": null,
			"Socializacion": null
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/respuesta_evaluacion/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRespuestaEvaluacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRespuestaEvaluacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRespuestaEvaluacion(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/respuesta_evaluacion/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRespuestaEvaluacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRespuestaEvaluacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
