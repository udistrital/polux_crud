package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllPreguntaFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/pregunta_formato/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllPreguntaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllPreguntaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllPreguntaFormato:", err.Error())
		t.Fail()
	}
}

func TestPostPreguntaFormato(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"Formato": {
		  "FechaRealizacion": "2023-10-08T19:00:00-05:00",
		  "Id": 3,
		  "Introduccion": "string",
		  "Nombre": "string"
		},
		"Id": 2,
		"Orden": 1,
		"Pregunta": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Enunciado": "string",
		  "Id": 3
		},
		"TipoPregunta": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 1,
		  "Nombre": "string"
		},
		"Valoracion": 2
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/pregunta_formato/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostPreguntaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostPreguntaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostPreguntaFormato:", err.Error())
		t.Fail()
	}
}

func TestGetOnePreguntaFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/pregunta_formato/2"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOnePreguntaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOnePreguntaFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOnePreguntaFormato:", err.Error())
		t.Fail()
	}
}

func TestPutPreguntaFormato(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"Formato": {
		  "FechaRealizacion": "2023-10-08T19:00:00-05:00",
		  "Id": 3,
		  "Introduccion": "string",
		  "Nombre": "string"
		},
		"Id": 2,
		"Orden": 1,
		"Pregunta": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Enunciado": "string",
		  "Id": 3
		},
		"TipoPregunta": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 1,
		  "Nombre": "string"
		},
		"Valoracion": 2
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/pregunta_formato/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutPreguntaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutPreguntaFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeletePreguntaFormato(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/pregunta_formato/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeletePreguntaFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeletePreguntaFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
