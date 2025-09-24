package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstructuraInvestigacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estructura_investigacion_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstructuraInvestigacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstructuraInvestigacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstructuraInvestigacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostEstructuraInvestigacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"CodigoEstructuraInvestigacion": 2,
		"Id": 1,
		"TrabajoGrado": {
		  "Id": 192
		}
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/estructura_investigacion_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstructuraInvestigacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstructuraInvestigacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstructuraInvestigacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstructuraInvestigacionTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estructura_investigacion_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstructuraInvestigacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstructuraInvestigacionTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstructuraInvestigacionTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutEstructuraInvestigacionTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"CodigoEstructuraInvestigacion": 3,
		"Id": 1,
		"TrabajoGrado": {
		  "Id": 192
		}
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estructura_investigacion_trabajo_grado/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstructuraInvestigacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstructuraInvestigacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstructuraInvestigacionTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estructura_investigacion_trabajo_grado/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstructuraInvestigacionTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstructuraInvestigacionTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
