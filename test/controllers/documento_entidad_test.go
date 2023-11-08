package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDocumentoEntidad(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_entidad/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDocumentoEntidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDocumentoEntidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDocumentoEntidad:", err.Error())
		t.Fail()
	}
}

func TestPostDocumentoEntidad(t *testing.T) {
	body := []byte(`{
		"DocumentoEscrito": {
		  "Enlace": "72107cb8-fac0-431f-bee0-9aa783d1c6e4",
		  "Id": 374,
		  "Resumen": "Acta de sustentación de el trabajo con id: 195, nombre:Lorem ipsum.",
		  "TipoDocumentoEscrito": 6,
		  "Titulo": "Acta de sustentación de trabajo id: 195"
		},
		"Entidad": 0,
		"Id": 1
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/documento_entidad/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDocumentoEntidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDocumentoEntidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDocumentoEntidad:", err.Error())
		t.Fail()
	}
}

func TestGetOneDocumentoEntidad(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_entidad/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDocumentoEntidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDocumentoEntidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDocumentoEntidad:", err.Error())
		t.Fail()
	}
}

func TestPutDocumentoEntidad(t *testing.T) {
	body := []byte(`{
		"DocumentoEscrito": {
		  "Enlace": "72107cb8-fac0-431f-bee0-9aa783d1c6e4",
		  "Id": 374,
		  "Resumen": "Acta de sustentación de el trabajo con id: 195, nombre:Lorem ipsum.",
		  "TipoDocumentoEscrito": 6,
		  "Titulo": "Acta de sustentación de trabajo id: 195"
		},
		"Entidad": 0,
		"Id": 1
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/documento_entidad/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDocumentoEntidad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDocumentoEntidad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDocumentoEntidad(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/documento_entidad/1", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDocumentoEntidad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDocumentoEntidad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
