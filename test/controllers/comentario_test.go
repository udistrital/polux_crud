package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllComentario(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/comentario/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllComentario Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllComentario Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllComentario:", err.Error())
		t.Fail()
	}
}

func TestGetOneComentario(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/comentario/43"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneComentario Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneComentario Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneAreaComentario:", err.Error())
		t.Fail()
	}
}

func TestPostComentario(t *testing.T) {
	body := []byte(`{
        "Id": 111,
        "Comentario": "Por favor descargue el documento de observaciones",
        "Fecha": "2023-02-08T12:52:44.82-05:00",
        "Autor": "ALEJANDRO GONZÁLEZ",
        "Correccion": {
            "Id": 36,
            "Observacion": "",
            "Pagina": 0,
            "RevisionTrabajoGrado": null,
            "Documento": false
        }
    }`)

	if response, err := http.Post("http://localhost:9002/v1/comentario/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostComentario Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostComentario Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostComentario:", err.Error())
		t.Fail()
	}
}

func TestPutComentario(t *testing.T) {
	body := []byte(`{
        "Id": 111,
        "Comentario": "COMENTARIO DE PRUEBA",
        "Fecha": "2023-02-08T12:52:44.82-05:00",
        "Autor": "ALEJANDRO GONZÁLEZ",
        "Correccion": {
            "Id": 36,
            "Observacion": "",
            "Pagina": 0,
            "RevisionTrabajoGrado": null,
            "Documento": false
        }
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/comentario/111", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutComentario Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutComentario Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteComentario(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/comentario/111", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteComentario Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteComentario Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
