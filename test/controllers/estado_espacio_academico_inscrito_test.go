package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllEstadoEspacioAcademicoInscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_espacio_academico_inscrito/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllEstadoEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllEstadoEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllEstadoEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestGetOneEstadoEspacioAcademicoInscrito(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/estado_espacio_academico_inscrito/1"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneEstadoEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneEstadoEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneEstadoEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestPostEstadoEspacioAcademicoInscrito(t *testing.T) {
	body := []byte(`{
        "Id": 4,
        "Nombre": "PRUEBA",
        "Descripcion": "ES UNA PRUEBA",
        "CodigoAbreviacion": "PRB",
        "Activo": true
    }`)

	if response, err := http.Post("http://localhost:9002/v1/estado_espacio_academico_inscrito/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostEstadoEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostEstadoEspacioAcademicoInscrito Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostEstadoEspacioAcademicoInscrito:", err.Error())
		t.Fail()
	}
}

func TestPutEstadoEspacioAcademicoInscrito(t *testing.T) {
	body := []byte(`{
        "Id": 4,
        "Nombre": "PRUEBA",
        "Descripcion": "ES UNA PRUEBA",
        "CodigoAbreviacion": "PRB",
        "Activo": true
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/estado_espacio_academico_inscrito/4", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutEstadoEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutEstadoEspacioAcademicoInscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteEstadoEspacioAcademicoInscrito(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/estado_espacio_academico_inscrito/4", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteEstadoEspacioAcademicoInscrito Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteEstadoEspacioAcademicoInscrito Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
