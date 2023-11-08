package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllSocializacion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/socializacion/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllSocializacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllSocializacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllSocializacion:", err.Error())
		t.Fail()
	}
}

func TestPostSocializacion(t *testing.T) {
	body := []byte(`{
		"Id": 1,
		"Fecha": "2023-10-10T19:00:00-05:00",
		"Lugar": 2,
		"TrabajoGrado": {
			"Id": 197,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/socializacion/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostSocializacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostSocializacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostSocializacion:", err.Error())
		t.Fail()
	}
}

func TestGetOneSocializacion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/socializacion/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneSocializacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneSocializacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneSocializacion:", err.Error())
		t.Fail()
	}
}

func TestPutSocializacion(t *testing.T) {
	body := []byte(`{
		"Id": 1,
		"Fecha": "2023-10-10T19:00:00-05:00",
		"Lugar": 3,
		"TrabajoGrado": {
			"Id": 197,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/socializacion/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutSocializacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutSocializacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteSocializacion(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/socializacion/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteSocializacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteSocializacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
