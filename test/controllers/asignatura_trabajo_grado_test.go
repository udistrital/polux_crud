package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetAllAsignaturaTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/asignatura_trabajo_grado/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetAllAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetAllAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetAllAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestGetOneAsignaturaTrabajoGrado(t *testing.T) {
	if response, err := http.Get("http://localhost:9002/v1/asignatura_trabajo_grado/370"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetOneAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetOneAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetOneAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPostAsignaturaTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 371,
        "CodigoAsignatura": 2,
        "Periodo": 3,
        "Anio": 2023,
        "Calificacion": 4.25,
        "TrabajoGrado": {
            "Id": 195,
            "Titulo": "Lorem ipsum",
            "Modalidad": {
                "Id": 4,
                "Nombre": "Monografía",
                "Descripcion": "Modalidad en la cual se realiza un ejercicio de aproximación y solución a un problema de investigación o de innovación en un campo de conocimiento",
                "CodigoAbreviacion": "MONO",
                "Activa": true
            },
            "EstadoTrabajoGrado": {
                "Id": 19,
                "Nombre": "Notificado a coordinación con calificación",
                "Descripcion": "Las calificaciones del jurado y director se notifican a la coordinación del pregrado",
                "CodigoAbreviacion": "NTF",
                "Activo": true
            },
            "DistincionTrabajoGrado": null,
            "PeriodoAcademico": "2023-3",
            "Objetivo": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
        },
        "EstadoAsignaturaTrabajoGrado": {
            "Id": 2,
            "Nombre": "Cursado",
            "Descripcion": "El estudiante ha cursado la asignatura para el trabajo de grado",
            "CodigoAbreviacion": "CDO",
            "Activo": true
        }
    }`)

	if response, err := http.Post("http://localhost:9002/v1/asignatura_trabajo_grado/", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 201 {
			t.Error("Error PostAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("PostAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error PostAsignaturaTrabajoGrado:", err.Error())
		t.Fail()
	}
}

func TestPutAsignaturaTrabajoGrado(t *testing.T) {
	body := []byte(`{
        "Id": 371,
        "CodigoAsignatura": 2,
        "Periodo": 3,
        "Anio": 2023,
        "Calificacion": 3.2,
        "TrabajoGrado": {
            "Id": 195,
            "Titulo": "Lorem ipsum",
            "Modalidad": {
                "Id": 4,
                "Nombre": "Monografía",
                "Descripcion": "Modalidad en la cual se realiza un ejercicio de aproximación y solución a un problema de investigación o de innovación en un campo de conocimiento",
                "CodigoAbreviacion": "MONO",
                "Activa": true
            },
            "EstadoTrabajoGrado": {
                "Id": 19,
                "Nombre": "Notificado a coordinación con calificación",
                "Descripcion": "Las calificaciones del jurado y director se notifican a la coordinación del pregrado",
                "CodigoAbreviacion": "NTF",
                "Activo": true
            },
            "DistincionTrabajoGrado": null,
            "PeriodoAcademico": "2023-3",
            "Objetivo": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
        },
        "EstadoAsignaturaTrabajoGrado": {
            "Id": 2,
            "Nombre": "Cursado",
            "Descripcion": "El estudiante ha cursado la asignatura para el trabajo de grado",
            "CodigoAbreviacion": "CDO",
            "Activo": true
        }
    }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9002/v1/asignatura_trabajo_grado/371", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error PutAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("PutAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDeleteAsignaturaTrabajoGrado(t *testing.T) {

	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9002/v1/asignatura_trabajo_grado/371", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error DeleteAsignaturaTrabajoGrado Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("DeleteAsignaturaTrabajoGrado Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}
