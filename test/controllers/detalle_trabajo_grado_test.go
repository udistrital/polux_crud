package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllDetalleTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllDetalleTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllDetalleTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllDetalleTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneDetalleTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/detalle_trabajo_grado/42"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneDetalleTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneDetalleTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneDetalleTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostDetalleTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Parametro": "4549",
        "Valor": "NOMBRE REVISTA",
        "TrabajoGrado": {
            "Id": 102,
            "Titulo": "NOMBRE ARTICULO ACADEMICO",
            "Modalidad": {
                "Id": 8,
                "Nombre": "Producción de artículo académico",
                "Descripcion": "Modalidad en la que el estudiante presenta evidencia de una publicación o aceptación de un artículo cientifico en revista homologada por el sistema de indexación nacional publindex de Colciencias ",
                "CodigoAbreviacion": "PACAD",
                "Activa": true
            },
            "EstadoTrabajoGrado": {
                "Id": 2,
                "Nombre": "Cancelado",
                "Descripcion": "El trabajo de grado ha sido cancelado",
                "CodigoAbreviacion": "CNC",
                "Activo": true
            },
            "DistincionTrabajoGrado": null,
            "PeriodoAcademico": "2023-3",
            "Objetivo": ""
        }
    }`)

	if response, err := http.Post("http://localhost:9002/v1/detalle_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostDetalleTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostDetalleTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostDetalleTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutDetalleTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 47,
        "Parametro": "4549",
        "Valor": "CAMBIO DE VALOR",
        "TrabajoGrado": {
            "Id": 102,
            "Titulo": "NOMBRE ARTICULO ACADEMICO",
            "Modalidad": {
                "Id": 8,
                "Nombre": "Producción de artículo académico",
                "Descripcion": "Modalidad en la que el estudiante presenta evidencia de una publicación o aceptación de un artículo cientifico en revista homologada por el sistema de indexación nacional publindex de Colciencias ",
                "CodigoAbreviacion": "PACAD",
                "Activa": true
            },
            "EstadoTrabajoGrado": {
                "Id": 2,
                "Nombre": "Cancelado",
                "Descripcion": "El trabajo de grado ha sido cancelado",
                "CodigoAbreviacion": "CNC",
                "Activo": true
            },
            "DistincionTrabajoGrado": null,
            "PeriodoAcademico": "2023-3",
            "Objetivo": ""
        }
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/detalle_trabajo_grado/47", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutDetalleTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutDetalleTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteDetalleTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/detalle_trabajo_grado/47", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteDetalleTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteDetalleTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
