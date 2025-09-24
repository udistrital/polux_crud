package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDocumentoTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDocumentoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDocumentoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDocumentoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneDocumentoTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_trabajo_grado/327"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDocumentoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDocumentoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDocumentoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostDocumentoTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 329,
		"TrabajoGrado": {
			"Id": 195,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"DocumentoEscrito": {
			"Id": 360,
			"Titulo": "",
			"Enlace": "",
			"Resumen": "",
			"TipoDocumentoEscrito": 0
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/documento_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDocumentoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDocumentoTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDocumentoTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutDocumentoTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 329,
		"TrabajoGrado": {
			"Id": 195,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		},
		"DocumentoEscrito": {
			"Id": 360,
			"Titulo": "",
			"Enlace": "",
			"Resumen": "",
			"TipoDocumentoEscrito": 0
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/documento_trabajo_grado/329", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDocumentoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDocumentoTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDocumentoTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/documento_trabajo_grado/329", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDocumentoTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDocumentoTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
