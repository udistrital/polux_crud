package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDocumentoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_solicitud/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDocumentoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDocumentoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDocumentoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestGetOneDocumentoSolicitud(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/documento_solicitud/537"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDocumentoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDocumentoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDocumentoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPostDocumentoSolicitud(t *testing.T) {
	body := []byte(`{
        "Id": 538,
        "DocumentoEscrito": {
            "Id": 94,
            "Titulo": "Acta 22 Codigo de carrera: 22 Fecha: 17-11-2022",
            "Enlace": "",
            "Resumen": "Acta de consejo de carrera del proyecto curricular",
            "TipoDocumentoEscrito": 1
        },
        "SolicitudTrabajoGrado": {
            "Id": 175,
            "Fecha": "2023-02-08T05:49:51.432-05:00",
            "ModalidadTipoSolicitud": {
                "Id": 20,
                "TipoSolicitud": {
                    "Id": 2,
                    "Nombre": "",
                    "Descripcion": "",
                    "CodigoAbreviacion": "",
                    "Activo": false
                },
                "Modalidad": {
                    "Id": 4,
                    "Nombre": "",
                    "Descripcion": "",
                    "CodigoAbreviacion": "",
                    "Activa": false
                }
            },
            "TrabajoGrado": null,
            "PeriodoAcademico": "2022-3"
        }
    }`)

	if response, err := http.Post("http://localhost:9002/v1/documento_solicitud/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDocumentoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDocumentoSolicitud Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDocumentoSolicitud:", err.Error())
		t.Fail()
	}
}

func TestPutDocumentoSolicitud(t *testing.T) {
	body := []byte(`{
        "Id": 538,
        "DocumentoEscrito": {
            "Id": 94,
            "Titulo": "Acta 22 Codigo de carrera: 22 Fecha: 17-11-2022",
            "Enlace": "",
            "Resumen": "Acta de consejo de carrera del proyecto curricular",
            "TipoDocumentoEscrito": 1
        },
        "SolicitudTrabajoGrado": {
            "Id": 175,
            "Fecha": "2023-02-08T05:49:51.432-05:00",
            "ModalidadTipoSolicitud": {
                "Id": 20,
                "TipoSolicitud": {
                    "Id": 2,
                    "Nombre": "",
                    "Descripcion": "",
                    "CodigoAbreviacion": "",
                    "Activo": false
                },
                "Modalidad": {
                    "Id": 4,
                    "Nombre": "",
                    "Descripcion": "",
                    "CodigoAbreviacion": "",
                    "Activa": false
                }
            },
            "TrabajoGrado": null,
            "PeriodoAcademico": "2022-3"
        }
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/documento_solicitud/538", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDocumentoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDocumentoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDocumentoSolicitud(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/documento_solicitud/538", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDocumentoSolicitud Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDocumentoSolicitud Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
