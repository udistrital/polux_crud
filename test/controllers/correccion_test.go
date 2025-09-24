package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllCorreccion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/correccion/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllCorreccion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllCorreccion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllCorreccion:", err.Error())
		t.Fail()
	}
}

func TestGetOneCorreccion(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/correccion/39"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneCorreccion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneCorreccion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneCorreccion:", err.Error())
		t.Fail()
	}
}

func TestPostCorreccion(t *testing.T) {
	body := []byte(`{
        "Id": 93,
        "Observacion": "PRUEBA",
        "Pagina": 1,
        "RevisionTrabajoGrado": {
            "Id": 30,
            "NumeroRevision": 0,
            "FechaRecepcion": "0001-01-01T00:00:00Z",
            "FechaRevision": "0001-01-01T00:00:00Z",
            "EstadoRevisionTrabajoGrado": null,
            "DocumentoTrabajoGrado": null,
            "VinculacionTrabajoGrado": null
        },
        "Documento": false
    }`)

	if response, err := http.Post("http://localhost:9002/v1/correccion/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostCorreccion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostCorreccion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostCorreccion:", err.Error())
		t.Fail()
	}
}

func TestPutCorreccion(t *testing.T) {
	body := []byte(`{
        "Id": 93,
        "Observacion": "PRUEBA DE PUT",
        "Pagina": 1,
        "RevisionTrabajoGrado": {
            "Id": 30,
            "NumeroRevision": 0,
            "FechaRecepcion": "0001-01-01T00:00:00Z",
            "FechaRevision": "0001-01-01T00:00:00Z",
            "EstadoRevisionTrabajoGrado": null,
            "DocumentoTrabajoGrado": null,
            "VinculacionTrabajoGrado": null
        },
        "Documento": false
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/correccion/93", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutCorreccion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutCorreccion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteCorreccion(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/correccion/93", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteCorreccion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteCorreccion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
