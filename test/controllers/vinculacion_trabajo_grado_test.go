package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllVinculacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/vinculacion_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllVinculacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllVinculacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllVinculacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneVinculacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/vinculacion_trabajo_grado/350"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneVinculacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneVinculacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneVinculacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostVinculacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 352,
		"Usuario": 1111111,
		"Activo": true,
		"FechaInicio": "2023-10-10T13:37:02.49-05:00",
		"FechaFin": "0001-01-01T00:00:00Z",
		"RolTrabajoGrado": {
			"Id": 1,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"TrabajoGrado": {
			"Id": 208,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/vinculacion_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostVinculacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostVinculacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostVinculacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutVinculacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 352,
		"Usuario": 2222222,
		"Activo": true,
		"FechaInicio": "2023-10-10T13:37:02.49-05:00",
		"FechaFin": "0001-01-01T00:00:00Z",
		"RolTrabajoGrado": {
			"Id": 1,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"TrabajoGrado": {
			"Id": 208,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/vinculacion_trabajo_grado/352", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutVinculacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutVinculacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteVinculacionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/vinculacion_trabajo_grado/352", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteVinculacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteVinculacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
