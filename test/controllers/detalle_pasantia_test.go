package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDetallePasantia(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_pasantia/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDetallePasantia Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDetallePasantia Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDetallePasantia:", err.Error())
		t.Fail()
	}
}

func TestGetOneDetallePasantia(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_pasantia/56"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDetallePasantia Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDetallePasantia Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDetallePasantia:", err.Error())
		t.Fail()
	}
}

func TestPostDetallePasantia(t *testing.T) {
	body := []byte(`{
		"Empresa": 0,
		"Horas": 0,
		"Id": 57,
		"ObjetoContrato": "Contrato de aprendizaje",
		"Observaciones": "PRUEBA",
		"TrabajoGrado": {
		  "DistincionTrabajoGrado": null,
		  "EstadoTrabajoGrado": null,
		  "Id": 188,
		  "Modalidad": null,
		  "Objetivo": "PRUEBA",
		  "PeriodoAcademico": "3",
		  "Titulo": "PRUEBA"
		}
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/detalle_pasantia/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDetallePasantia Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDetallePasantia Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDetallePasantia:", err.Error())
		t.Fail()
	}
}

func TestPutDetallePasantia(t *testing.T) {
	body := []byte(`{
		"Empresa": 0,
		"Horas": 0,
		"Id": 57,
		"ObjetoContrato": "Contrato de aprendizaje",
		"Observaciones": "PRUEBA",
		"TrabajoGrado": {
		  "DistincionTrabajoGrado": null,
		  "EstadoTrabajoGrado": null,
		  "Id": 188,
		  "Modalidad": null,
		  "Objetivo": "PRUEBA",
		  "PeriodoAcademico": "3",
		  "Titulo": "PRUEBA"
		}
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/detalle_pasantia/57", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDetallePasantia Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDetallePasantia Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDetallePasantia(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/detalle_pasantia/57", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDetallePasantia Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDetallePasantia Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
