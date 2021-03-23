package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	_ "github.com/udistrital/polux_crud/routers"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
}

// TestGet is a sample to run an endpoint test
func TestGetAllObject(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/modalidad", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetOneObject(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/modalidad/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOneObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

type Modalidad struct {
	Id          int    `orm:"column(id);pk;auto"`
	Nombre      string `orm:"column(nombre)"`
	Activa      bool   `orm:"column(activa)"`
	Descripcion string `orm:"column(descripcion);null"`
}

// TestGet is a sample to run an endpoint test
func TestPostObject(t *testing.T) {
	// Create an instance of the Object struct.
	object := Modalidad{
		Id:          11,
		Nombre:      "Modalidad de prueba",
		Activa:      true,
		Descripcion: "Descripción de la modalidad",
	}

	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(object)
	// Convert bytes to string.
	s := string(b)
	// Convert string to io.Reader
	bu := bytes.NewBuffer([]byte(s))

	r, _ := http.NewRequest("POST", "/v1/modalidad", bu)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPostObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPutObject(t *testing.T) {
	object := Modalidad{
		Id:          1,
		Nombre:      "Monografía",
		Activa:      true,
		Descripcion: "Modalidad de Monografía",
	}
	b, _ := json.Marshal(object)
	s := string(b)
	bu := bytes.NewBuffer([]byte(s))

	r, _ := http.NewRequest("PUT", "/v1/modalidad/1", bu)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPutObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
