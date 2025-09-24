package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDetalleSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDetalleSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDetalleSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDetalleSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneDetalleSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_solicitud/4269"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDetalleSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDetalleSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDetalleSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostDetalleSolicitud(t *testing.T) {
	body := []byte(`{
		"Descripcion": "PRUEBA",
		"DetalleTipoSolicitud": {
			"Id": 252,
			"Detalle": null,
			"ModalidadTipoSolicitud": null,
			"Activo": false,
			"Requerido": false,
			"NumeroOrden": 0
		},
		"Id": 4270,
		"SolicitudTrabajoGrado": {
			"Id": 664,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		}
	  }`)

	if response, err := http.Post("http://localhost:9002/v1/detalle_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDetalleSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDetalleSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDetalleSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutDetalleSolicitud(t *testing.T) {
	body := []byte(`{
		"Descripcion": "CAMBIO_PRUEBA",
		"DetalleTipoSolicitud": {
			"Id": 252,
			"Detalle": null,
			"ModalidadTipoSolicitud": null,
			"Activo": false,
			"Requerido": false,
			"NumeroOrden": 0
		},
		"Id": 4270,
		"SolicitudTrabajoGrado": {
			"Id": 664,
			"Fecha": "0001-01-01T00:00:00Z",
			"ModalidadTipoSolicitud": null,
			"TrabajoGrado": null,
			"PeriodoAcademico": ""
		}
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/detalle_solicitud/4270", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDetalleSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDetalleSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDetalleSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/detalle_solicitud/4270", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDetalleSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDetalleSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
