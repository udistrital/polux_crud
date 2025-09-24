package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRespuesta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRespuesta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRespuesta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRespuesta:", err.Error())
		t.Fail()
	}
}

func TestGetOneRespuesta(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRespuesta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRespuesta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRespuesta:", err.Error())
		t.Fail()
	}
}

func TestPostRespuesta(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Descripcion": "PRUEBA",
		"Nombre": "PRUEBA",
		"CodigoAbreviacion": "PR",
		"Activo": true
	}`)

	if response, err := http.Post("http://localhost:9002/v1/respuesta/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRespuesta Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRespuesta Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRespuesta:", err.Error())
		t.Fail()
	}
}

func TestPutRespuesta(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Descripcion": "PRUEBA",
		"Nombre": "PRUEBA",
		"CodigoAbreviacion": "PRR",
		"Activo": true
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/respuesta/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRespuesta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRespuesta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRespuesta(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/respuesta/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRespuesta Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRespuesta Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
