package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllTipoDetalle(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_detalle/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllTipoDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllTipoDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllTipoDetalle:", err.Error())
		t.Fail()
	}
}

func TestGetOneTipoDetalle(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/tipo_detalle/10"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneTipoDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneTipoDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneTipoDetalle:", err.Error())
		t.Fail()
	}
}

func TestPostTipoDetalle(t *testing.T) {
	body := []byte(`{
		"Id": 12,
		"Nombre": "PRUEBA",
		"Descripcion": "",
		"CodigoAbreviacion": "lbl",
		"Activo": true
	}`)

	if response, err := http.Post("http://localhost:9002/v1/tipo_detalle/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostTipoDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostTipoDetalle Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostTipoDetalle:", err.Error())
		t.Fail()
	}
}

func TestPutTipoDetalle(t *testing.T) {
	body := []byte(`{
		"Id": 12,
		"Nombre": "PRUEBA",
		"Descripcion": "",
		"CodigoAbreviacion": "lbl",
		"Activo": false
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/tipo_detalle/12", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutTipoDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutTipoDetalle Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteTipoDetalle(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/tipo_detalle/12", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteTipoDetalle Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteTipoDetalle Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
