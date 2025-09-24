package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDetalle(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDetalle:", err.Error())
		t.Fail()
	}
}

func TestGetOneDetalle(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle/74"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDetalle:", err.Error())
		t.Fail()
	}
}

func TestPostDetalle(t *testing.T) {
	body := []byte(`{
		"Id": 76,
		"Nombre": "PRUEBA",
		"Enunciado": "PRUEBA",
		"Descripcion": "no_service",
		"CodigoAbreviacion": "PRUEBA",
		"Activo": true,
		"TipoDetalle": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		}
	}`)

	if response, err := http.Post("http://localhost:9002/v1/detalle/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDetalle:", err.Error())
		t.Fail()
	}
}

func TestPutDetalle(t *testing.T) {
	body := []byte(`{
		"Id": 76,
		"Nombre": "CAMBIO PRUEBA",
		"Enunciado": "PRUEBA",
		"Descripcion": "no_service",
		"CodigoAbreviacion": "PRUEBA",
		"Activo": true,
		"TipoDetalle": {
			"Id": 2,
			"Nombre": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false
		}
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/detalle/76", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDetalle Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDetalle(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/detalle/76", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDetalle Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
