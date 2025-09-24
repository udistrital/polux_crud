package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRevisionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/revision_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneRevisionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/revision_trabajo_grado/70"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostRevisionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 73,
		"NumeroRevision": 2,
		"FechaRecepcion": "2023-10-02T10:34:24.228-05:00",
		"FechaRevision": "2023-10-02T10:34:48.769-05:00",
		"EstadoRevisionTrabajoGrado": {
			"Id": 3,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"DocumentoTrabajoGrado": {
			"Id": 325,
			"TrabajoGrado": null,
			"DocumentoEscrito": null
		},
		"VinculacionTrabajoGrado": {
			"Id": 321,
			"Usuario": 0,
			"Activo": false,
			"FechaInicio": "0001-01-01T00:00:00Z",
			"FechaFin": "0001-01-01T00:00:00Z",
			"RolTrabajoGrado": null,
			"TrabajoGrado": null
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/revision_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRevisionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRevisionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutRevisionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 73,
		"NumeroRevision": 2,
		"FechaRecepcion": "2023-10-02T10:34:24.228-05:00",
		"FechaRevision": "2023-10-02T10:34:48.769-05:00",
		"EstadoRevisionTrabajoGrado": {
			"Id": 3,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"DocumentoTrabajoGrado": {
			"Id": 325,
			"TrabajoGrado": null,
			"DocumentoEscrito": null
		},
		"VinculacionTrabajoGrado": {
			"Id": 321,
			"Usuario": 0,
			"Activo": false,
			"FechaInicio": "0001-01-01T00:00:00Z",
			"FechaFin": "0001-01-01T00:00:00Z",
			"RolTrabajoGrado": null,
			"TrabajoGrado": null
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/revision_trabajo_grado/73", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRevisionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRevisionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/revision_trabajo_grado/73", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRevisionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRevisionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
