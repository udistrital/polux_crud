package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllModalidadTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/modalidad_tipo_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllModalidadTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllModalidadTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllModalidadTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneModalidadTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/modalidad_tipo_solicitud/90"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneModalidadTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneModalidadTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneModalidadTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostModalidadTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 95,
		"Modalidad": {
		  "Activa": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 9,
		  "Nombre": "string"
		},
		"TipoSolicitud": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 9,
		  "Nombre": "string"
		}
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/modalidad_tipo_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostModalidadTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostModalidadTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostModalidadTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutModalidadTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 95,
		"Modalidad": {
		  "Activa": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 9,
		  "Nombre": "string"
		},
		"TipoSolicitud": {
		  "Activo": true,
		  "CodigoAbreviacion": "string",
		  "Descripcion": "string",
		  "Id": 9,
		  "Nombre": "string"
		}
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/modalidad_tipo_solicitud/95", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutModalidadTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutModalidadTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteModalidadTipoSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/modalidad_tipo_solicitud/95", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteModalidadTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteModalidadTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
