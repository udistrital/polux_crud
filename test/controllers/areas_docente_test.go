package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPostAreasDocente(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"AreaConocimiento": {
		  "Activo": true,
		  "CodigoAbreviacion": "",
		  "Descripcion": "SubArea enfocada en el desarrollo y gestión de sistemas con bases en la ingenieria",
		  "Id": 1,
		  "Nombre": "Ingeniería de Sistemas",
		  "SniesArea": 7
		},
		"Docente": 0,
		"Id": 1
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/areas_docente/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostAreasDocente Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostAreasDocente Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostAreasDocente:", err.Error())
		t.Fail()
	}
}

func TestGetAllAreasDocente(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/areas_docente/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllAreasDocente Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllAreasDocente Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllAreasDocente:", err.Error())
		t.Fail()
	}
}

func TestGetOneAreasDocente(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/areas_docente/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneAreasDocente Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneAreasDocente Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneAreasDocente:", err.Error())
		t.Fail()
	}
}

func TestPutAreasDocente(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"AreaConocimiento": {
		  "Activo": true,
		  "CodigoAbreviacion": "",
		  "Descripcion": "SubArea enfocada en el desarrollo y gestión de sistemas con bases en la ingenieria",
		  "Id": 1,
		  "Nombre": "Ingeniería de Sistemas",
		  "SniesArea": 7
		},
		"Docente": 1234,
		"Id": 1
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/areas_docente/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutAreasDocente Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutAreasDocente Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteAreasDocente(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/areas_docente/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteAreasDocente Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteAreasDocente Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
