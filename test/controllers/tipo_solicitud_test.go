package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_solicitud/15"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 16,
		"Nombre": "PRUEBA",
		"Descripcion": "PRUEBA",
		"CodigoAbreviacion": "prb",
		"Activo": true
	}`)

	if response, err := http.Post("http://localhost:9002/v1/tipo_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 16,
		"Nombre": "PRUEBA",
		"Descripcion": "PRUEBA",
		"CodigoAbreviacion": "prb",
		"Activo": true
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/tipo_solicitud/16", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteTipoSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/tipo_solicitud/16", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
