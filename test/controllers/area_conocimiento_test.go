package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllAreaConocimiento(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/area_conocimiento/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllAreaConocimiento Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllAreaConocimiento Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllAreaConocimiento:", err.Error())
		t.Fail()
	}
}

func TestGetOneAreaConocimiento(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/area_conocimiento/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneAreaConocimiento Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneAreaConocimiento Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneAreaConocimiento:", err.Error())
		t.Fail()
	}
}

func TestPostAreaConocimiento(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "PRB",
		"Descripcion": "Area de conocimiento de prueba.",
		"Id": 3,
		"Nombre": "PRUEBA",
		"SniesArea": 7
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/area_conocimiento/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostAreaConocimiento Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostAreaConocimiento Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostAreaConocimiento:", err.Error())
		t.Fail()
	}
}

func TestPutAreaConocimiento(t *testing.T) {
	body := []byte(`{
		"Activo": true,
		"CodigoAbreviacion": "CAMBIO_PRUEBA",
		"Descripcion": "Area de conocimiento de prueba.",
		"Id": 3,
		"Nombre": "PRUEBA",
		"SniesArea": 7
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/area_conocimiento/3", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutAreaConocimiento Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutAreaConocimiento Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteAreaConocimiento(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/area_conocimiento/3", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteAreaConocimiento Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteAreaConocimiento Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
