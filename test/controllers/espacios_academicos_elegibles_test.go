package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEspaciosAcademicosElegibles(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/espacios_academicos_elegibles/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEspaciosAcademicosElegibles Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEspaciosAcademicosElegibles Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEspaciosAcademicosElegibles:", err.Error())
		t.Fail()
	}
}

func TestGetOneEspaciosAcademicosElegibles(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/espacios_academicos_elegibles/48"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEspaciosAcademicosElegibles Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEspaciosAcademicosElegibles Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEspaciosAcademicosElegibles:", err.Error())
		t.Fail()
	}
}

func TestPostEspaciosAcademicosElegibles(t *testing.T) {
	body := []byte(`{
		"Id": 50,
		"CodigoAsignatura": 1111111,
		"Activo": true,
		"CarreraElegible": {
			"Id": 23,
			"CodigoCarrera": 0,
			"CuposExcelencia": 0,
			"CuposAdicionales": 0,
			"Periodo": 0,
			"Anio": 0,
			"CodigoPensum": 0,
			"Nivel": ""
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/espacios_academicos_elegibles/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEspaciosAcademicosElegibles Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEspaciosAcademicosElegibles Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEspaciosAcademicosElegibles:", err.Error())
		t.Fail()
	}
}

func TestPutEspaciosAcademicosElegibles(t *testing.T) {
	body := []byte(`{
		"Id": 50,
		"CodigoAsignatura": 2222222,
		"Activo": true,
		"CarreraElegible": {
			"Id": 23,
			"CodigoCarrera": 0,
			"CuposExcelencia": 0,
			"CuposAdicionales": 0,
			"Periodo": 0,
			"Anio": 0,
			"CodigoPensum": 0,
			"Nivel": ""
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/espacios_academicos_elegibles/50", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEspaciosAcademicosElegibles Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEspaciosAcademicosElegibles Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEspaciosAcademicosElegibles(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/espacios_academicos_elegibles/50", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEspaciosAcademicosElegibles Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEspaciosAcademicosElegibles Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
