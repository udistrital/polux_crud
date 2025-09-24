package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRespuestaFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_formato/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRespuestaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRespuestaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRespuestaFormato:", err.Error())
		t.Fail()
	}
}

func TestGetOneRespuestaFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_formato/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRespuestaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRespuestaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRespuestaFormato:", err.Error())
		t.Fail()
	}
}

func TestPostRespuestaFormato(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Orden": 1,
		"Valoracion": 3,
		"Respuesta": {
			"Id": 3,
			"Descripcion": "",
			"Nombre": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"PreguntaFormato": {
			"Id": 1,
			"Activo": false,
			"Orden": 0,
			"Valoracion": 0,
			"Pregunta": null,
			"Formato": null,
			"TipoPregunta": null
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/respuesta_formato/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRespuestaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRespuestaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRespuestaFormato:", err.Error())
		t.Fail()
	}
}

func TestPutRespuestaFormato(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Orden": 1,
		"Valoracion": 3,
		"Respuesta": {
			"Id": 3,
			"Descripcion": "",
			"Nombre": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"PreguntaFormato": {
			"Id": 1,
			"Activo": false,
			"Orden": 0,
			"Valoracion": 0,
			"Pregunta": null,
			"Formato": null,
			"TipoPregunta": null
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/respuesta_formato/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRespuestaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRespuestaFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRespuestaFormato(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/respuesta_formato/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRespuestaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRespuestaFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
