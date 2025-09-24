package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllRolTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/rol_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllRolTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllRolTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllRolTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneRolTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/rol_trabajo_grado/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneRolTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneRolTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneRolTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostRolTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 6,
		"Nombre": "Director Interno",
		"Descripcion": "",
		"CodigoAbreviacion": "DIRECTOR",
		"Activo": true
	}`)

	if response, err := http.Post("http://localhost:9002/v1/rol_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostRolTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostRolTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostRolTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutRolTrabajoGrado(t *testing.T) {
	body := []byte(`{
		"Id": 6,
		"Nombre": "Director Interno Cambio",
		"Descripcion": "",
		"CodigoAbreviacion": "DIRECTOR",
		"Activo": true
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/rol_trabajo_grado/6", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutRolTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutRolTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteRolTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/rol_trabajo_grado/6", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteRolTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteRolTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
