package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllTipoPregunta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_pregunta/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllTipoPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllTipoPregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllTipoPregunta:", err.Error())
		t.Fail()
	}
}

func TestGetOneTipoPregunta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_pregunta/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneTipoPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneTipoPregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneTipoPregunta:", err.Error())
		t.Fail()
	}
}

func TestPostTipoPregunta(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Nombre": "PRUEBA2",
		"Descripcion": "Es una PRUEBA2",
		"CodigoAbreviacion": "PRUEBA2",
		"Activo": false
	}`)

	if response, err := http.Post("http://localhost:9002/v1/tipo_pregunta/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostTipoPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostTipoPregunta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostTipoPregunta:", err.Error())
		t.Fail()
	}
}

func TestPutTipoPregunta(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Nombre": "CAMBIO PRUEBA2",
		"Descripcion": "Es una PRUEBA2",
		"CodigoAbreviacion": "PRUEBA2",
		"Activo": false
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/tipo_pregunta/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutTipoPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutTipoPregunta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteTipoPregunta(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/tipo_pregunta/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteTipoPregunta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteTipoPregunta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
