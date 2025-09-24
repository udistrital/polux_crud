package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_solicitud/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoSolicitud(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 24,
		"Nombre": "PRUEBA"
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/estado_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoSolicitud(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Es una prueba",
		"Id": 24,
		"Nombre": "CAMBIO PRUEBA"
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_solicitud/24", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_solicitud/24", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
