package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/formato/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllFormato:", err.Error())
		t.Fail()
	}
}

func TestGetOneFormato(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/formato/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneFormato:", err.Error())
		t.Fail()
	}
}

func TestPostFormato(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Nombre": "PRUEBA2",
		"Introduccion": "INTRO PRUEBA2",
		"FechaRealizacion": "2023-10-08T19:00:00-05:00"
	}`)

	if response, err := http.Post("http://localhost:9002/v1/formato/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostFormato:", err.Error())
		t.Fail()
	}
}

func TestPutFormato(t *testing.T) {
	body := []byte(`{
		"Id": 2,
		"Nombre": "CAMBIO PRUEBA2",
		"Introduccion": "INTRO PRUEBA2",
		"FechaRealizacion": "2023-10-08T19:00:00-05:00"
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/formato/2", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteFormato(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/formato/2", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteFormato Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteFormato Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
