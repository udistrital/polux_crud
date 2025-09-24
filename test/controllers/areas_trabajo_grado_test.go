package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllAreasTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/areas_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllAreasTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllAreasTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllAreasTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneAreasTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/areas_trabajo_grado/161"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneAreasTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneAreasTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneAreasTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostAreasTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 166,
        "AreaConocimiento": {
            "Id": 2,
            "Nombre": "Ingeniería de Sistemas",
            "Descripcion": "SubArea enfocada en el desarrollo y gestión de sistemas con bases en la ingenieria",
            "CodigoAbreviacion": "",
            "Activo": true,
            "SniesArea": 7
        },
        "TrabajoGrado": {
            "Id": 182,
            "Titulo": "MONOGRAFÍA DANIEL",
            "Modalidad": {
                "Id": 4,
                "Nombre": "Monografía",
                "Descripcion": "Modalidad en la cual se realiza un ejercicio de aproximación y solución a un problema de investigación o de innovación en un campo de conocimiento",
                "CodigoAbreviacion": "MONO",
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
            "Objetivo": "OBJETIVOS"
        },
        "Activo": true
    }`)

	if response, err := http.Post("http://localhost:9002/v1/areas_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostAreasTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostAreasTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostAreasTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutAreasTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 166,
        "AreaConocimiento": {
            "Id": 2,
            "Nombre": "Ingeniería de Sistemas",
            "Descripcion": "SubArea enfocada en el desarrollo y gestión de sistemas con bases en la ingenieria",
            "CodigoAbreviacion": "",
            "Activo": true,
            "SniesArea": 7
        },
        "TrabajoGrado": {
            "Id": 182,
            "Titulo": "MONOGRAFÍA DANIEL",
            "Modalidad": {
                "Id": 4,
                "Nombre": "Monografía",
                "Descripcion": "Modalidad en la cual se realiza un ejercicio de aproximación y solución a un problema de investigación o de innovación en un campo de conocimiento",
                "CodigoAbreviacion": "MONO",
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
            "Objetivo": "OBJETIVOS"
        },
        "Activo": true
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/areas_trabajo_grado/166", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutAreasTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutAreasTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteAreasTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/areas_trabajo_grado/166", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteAreasTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteAreasTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
