package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllSolicitudTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/solicitud_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllSolicitudTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllSolicitudTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllSolicitudTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneSolicitudTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/solicitud_trabajo_grado/678"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneSolicitudTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneSolicitudTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneSolicitudTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostSolicitudTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 680,
		"Fecha": "2023-10-10T08:23:50.696-05:00",
		"ModalidadTipoSolicitud": {
			"Id": 22,
			"TipoSolicitud": null,
			"Modalidad": null
		},
		"TrabajoGrado": {
			"Id": 208,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"PeriodoAcademico": "2023-3"
	}`)

	if response, err := http.Post("http://localhost:9002/v1/solicitud_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostSolicitudTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostSolicitudTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostSolicitudTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutSolicitudTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 680,
		"Fecha": "2023-10-10T08:23:50.696-05:00",
		"ModalidadTipoSolicitud": {
			"Id": 22,
			"TipoSolicitud": null,
			"Modalidad": null
		},
		"TrabajoGrado": {
			"Id": 208,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"PeriodoAcademico": "2023-1"
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/solicitud_trabajo_grado/680", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutSolicitudTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutSolicitudTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteSolicitudTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/solicitud_trabajo_grado/680", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteSolicitudTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteSolicitudTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
