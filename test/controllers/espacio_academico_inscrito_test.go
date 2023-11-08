package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEspacioAcademicoInscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/espacio_academico_inscrito/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestGetOneEspacioAcademicoInscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/espacio_academico_inscrito/45"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestPostEspacioAcademicoInscrito(t *testing.T) {
	body := []byte(`{
		"Id": 47,
		"Nota": 4.5,
		"EspaciosAcademicosElegibles": {
			"Id": 30,
			"CodigoAsignatura": 0,
			"Activo": false,
			"CarreraElegible": null
		},
		"EstadoEspacioAcademicoInscrito": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"TrabajoGrado": {
			"Id": 175,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/espacio_academico_inscrito/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestPutEspacioAcademicoInscrito(t *testing.T) {
	body := []byte(`{
		"Id": 47,
		"Nota": 4.5,
		"EspaciosAcademicosElegibles": {
			"Id": 30,
			"CodigoAsignatura": 0,
			"Activo": false,
			"CarreraElegible": null
		},
		"EstadoEspacioAcademicoInscrito": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		},
		"TrabajoGrado": {
			"Id": 175,
			"Titulo": "",
			"Modalidad": null,
			"EstadoTrabajoGrado": null,
			"DistincionTrabajoGrado": null,
			"PeriodoAcademico": "",
			"Objetivo": ""
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/espacio_academico_inscrito/47", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEspacioAcademicoInscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEspacioAcademicoInscrito(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/espacio_academico_inscrito/47", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEspacioAcademicoInscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
