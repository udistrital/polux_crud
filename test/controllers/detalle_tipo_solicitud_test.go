package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDetalleTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_tipo_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDetalleTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDetalleTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDetalleTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneDetalleTipoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_tipo_solicitud/311"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDetalleTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDetalleTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDetalleTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostDetalleTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 312,
		"Detalle": {
			"Id": 1,
			"Nombre": "",
			"Enunciado": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false,
			"TipoDetalle": null
		},
		"ModalidadTipoSolicitud": {
			"Id": 55,
			"TipoSolicitud": null,
			"Modalidad": null
		},
		"Activo": true,
		"Requerido": true,
		"NumeroOrden": 11
	}`)

	if response, err := http.Post("http://localhost:9002/v1/detalle_tipo_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDetalleTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDetalleTipoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDetalleTipoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutDetalleTipoSolicitud(t *testing.T) {
	body := []byte(`{
		"Id": 312,
		"Detalle": {
			"Id": 1,
			"Nombre": "",
			"Enunciado": "",
			"Descripcion": "",
			"CodigoAbreviacion": "",
			"Activo": false,
			"TipoDetalle": null
		},
		"ModalidadTipoSolicitud": {
			"Id": 55,
			"TipoSolicitud": null,
			"Modalidad": null
		},
		"Activo": false,
		"Requerido": true,
		"NumeroOrden": 11
	}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/detalle_tipo_solicitud/312", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDetalleTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDetalleTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDetalleTipoSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/detalle_tipo_solicitud/312", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDetalleTipoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDetalleTipoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
