package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDocumentoEscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_escrito/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDocumentoEscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDocumentoEscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDocumentoEscrito:", err.Error())
		t.Fail()
	}
}

func TestGetOneDocumentoEscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_escrito/374"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDocumentoEscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDocumentoEscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDocumentoEscrito:", err.Error())
		t.Fail()
	}
}

func TestPostDocumentoEscrito(t *testing.T) {
	body := []byte(`{
		"Id": 375,
		"Titulo": "PRUEBA",
		"Enlace": "72107cb8-fac0-431f-bee0-9aa783d1c6e4",
		"Resumen": "RESUMEN DE PRUEBA",
		"TipoDocumentoEscrito": 6
	}`)

	if response, err := http.Post("http://localhost:9002/v1/documento_escrito/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDocumentoEscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDocumentoEscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDocumentoEscrito:", err.Error())
		t.Fail()
	}
}

func TestPutDocumentoEscrito(t *testing.T) {
	body := []byte(`{
		"Id": 375,
		"Titulo": "CAMBIO PRUEBA",
		"Enlace": "72107cb8-fac0-431f-bee0-9aa783d1c6e4",
		"Resumen": "RESUMEN DE PRUEBA",
		"TipoDocumentoEscrito": 6
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/documento_escrito/375", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDocumentoEscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDocumentoEscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDocumentoEscrito(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/documento_escrito/375", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDocumentoEscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDocumentoEscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
