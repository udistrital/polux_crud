package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllModalidad(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/modalidad/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllModalidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllModalidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllModalidad:", err.Error())
		t.Fail()
	}
}

func TestGetOneModalidad(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/modalidad/9"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneModalidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneModalidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneModalidad:", err.Error())
		t.Fail()
	}
}

func TestPostModalidad(t *testing.T) {
	body := []byte(`{
		"Activa": true,
		"CodigoAbreviacion": "PRUEBA",
		"Descripcion": "Es una prueba",
		"Id": 10,
		"Nombre": "PRUEBA"
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/modalidad/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostModalidad Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostModalidad Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostModalidad:", err.Error())
		t.Fail()
	}
}

func TestPutModalidad(t *testing.T) {
	body := []byte(`{
		"Activa": true,
		"CodigoAbreviacion": "PRUEBA",
		"Descripcion": "Es una prueba",
		"Id": 10,
		"Nombre": "PRUEBA"
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/modalidad/10", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutModalidad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutModalidad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteModalidad(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/modalidad/10", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteModalidad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteModalidad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
