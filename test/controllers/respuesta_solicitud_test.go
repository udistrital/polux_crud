package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRespuestaSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRespuestaSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRespuestaSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRespuestaSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneRespuestaSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/respuesta_solicitud/1265"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRespuestaSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRespuestaSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRespuestaSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostRespuestaSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 1268,
		"Fecha": "2023-10-10T14:20:28.694-05:00",
		"Justificacion": "El Director aprobo la Solicitud de revisión de trabajo de grado",
		"EnteResponsable": 0,
		"Usuario": 1111111,
		"EstadoSolicitud": {
			"Id": 17,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"SolicitudTrabajoGrado": {
			"Id": 648,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		},
		"Activo": false
	}`)

	if response, err := http.Post("http://localhost:9002/v1/respuesta_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRespuestaSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRespuestaSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRespuestaSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutRespuestaSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 1268,
		"Fecha": "2023-10-10T14:20:28.694-05:00",
		"Justificacion": "El Director aprobo la Solicitud de revisión de trabajo de grado",
		"EnteResponsable": 0,
		"Usuario": 1111111,
		"EstadoSolicitud": {
			"Id": 17,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"SolicitudTrabajoGrado": {
			"Id": 648,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		},
		"Activo": false
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/respuesta_solicitud/1268", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRespuestaSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRespuestaSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRespuestaSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/respuesta_solicitud/1268", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRespuestaSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRespuestaSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
