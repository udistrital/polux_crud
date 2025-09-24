package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllPregunta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/pregunta/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllPregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllPregunta:", err.Error())
		t.Fail()
	}
}

func TestGetOnePregunta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/pregunta/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOnePregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOnePregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOnePregunta:", err.Error())
		t.Fail()
	}
}

func TestPostPregunta(t *testing.T) {
	body := []byte(`{
		"Activo": false,
		"CodigoAbreviacion": "PRUEBA2",
		"Descripcion": "PRUEBA2",
		"Enunciado": "PRUEBA2",
		"Id": 2
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/pregunta/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostPregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostPregunta:", err.Error())
		t.Fail()
	}
}

func TestPutPregunta(t *testing.T) {
	body := []byte(`{
		"Activo": false,
		"CodigoAbreviacion": "PRUEBA2",
		"Descripcion": "PRUEBA2",
		"Enunciado": "CAMBIO PRUEBA2",
		"Id": 2
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/pregunta/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutPregunta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeletePregunta(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/pregunta/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeletePregunta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeletePregunta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
