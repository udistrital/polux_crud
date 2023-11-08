package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllCarreraElegible(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/carrera_elegible/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllCarreraElegible Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllCarreraElegible Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllCarreraElegible:", err.Error())
		t.Fail()
	}
}

func TestGetOneCarreraElegible(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/carrera_elegible/23"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneCarreraElegible Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneCarreraElegible Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneCarreraElegible:", err.Error())
		t.Fail()
	}
}

func TestPostCarreraElegible(t *testing.T) {
	body := []byte(`{
        "Id": 24,
        "CodigoCarrera": 199,
        "CuposExcelencia": 10,
        "CuposAdicionales": 10,
        "Periodo": 1,
        "Anio": 2025,
        "CodigoPensum": 3,
        "Nivel": "POSGRADO"
    }`)

	if response, err := http.Post("http://localhost:9002/v1/carrera_elegible/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostCarreraElegible Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostCarreraElegible Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostCarreraElegible:", err.Error())
		t.Fail()
	}
}

func TestPutCarreraElegible(t *testing.T) {
	body := []byte(`{
        "Id": 24,
        "CodigoCarrera": 199,
        "CuposExcelencia": 10,
        "CuposAdicionales": 10,
        "Periodo": 3,
        "Anio": 2025,
        "CodigoPensum": 3,
        "Nivel": "POSGRADO"
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/carrera_elegible/24", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutCarreraElegible Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutCarreraElegible Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteCarreraElegible(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/carrera_elegible/24", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteCarreraElegible Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteCarreraElegible Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
