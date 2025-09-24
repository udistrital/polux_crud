package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllUsuarioSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/usuario_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllUsuarioSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllUsuarioSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllUsuarioSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneUsuarioSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/usuario_solicitud/695"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneUsuarioSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneUsuarioSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneUsuarioSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostUsuarioSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 696,
		"Usuario": "20161778011",
		"SolicitudTrabajoGrado": {
			"Id": 679,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/usuario_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostUsuarioSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostUsuarioSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostUsuarioSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutUsuarioSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 696,
		"Usuario": "111111",
		"SolicitudTrabajoGrado": {
			"Id": 679,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/usuario_solicitud/696", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutUsuarioSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutUsuarioSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteUsuarioSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/usuario_solicitud/696", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteUsuarioSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteUsuarioSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
