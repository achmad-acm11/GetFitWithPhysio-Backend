package patient

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/getfitwith_physio?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	helper.HandleError(err)

	return db
}

func setupRouter(db *gorm.DB) *httprouter.Router {
	// Init Validate
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	router = Config(db, validate, router)
	router.PanicHandler = exception.ErrorHandler

	return router
}
func truncatePatient(db *gorm.DB) {
	db.Exec("TRUNCATE patients")
}

// func truncateUser(db *gorm.DB) {
// 	db.Exec("TRUNCATE users")
// }
func TestGetAllPatientSuccess(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/patients", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestRegisterPatientSuccess(t *testing.T) {
	db := setupDatabase()
	truncatePatient(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": "patient1",
	"gender": "Perempuan",
	"phone": "0895619258715",
	"address": "link harum manis rt04/02 no 166 kel cirimekar kec.cibinong Kab.Bogor",
	"nik": "3215646516543",
	"birthdate": "15-09-1997",
	"email": "patient@gmail.com",
	"password": "12345678"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register", reqBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
	assert.Equal(t, "patient1", responseBody["data"].(map[string]interface{})["name"])
}
func TestRegisterPatientFailed(t *testing.T) {
	db := setupDatabase()
	truncatePatient(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"name": "",
	"gender": "Perempuan",
	"phone": "0895619258715",
	"address": "link harum manis rt04/02 no 166 kel cirimekar kec.cibinong Kab.Bogor",
	"nik": "3215646516543",
	"birthdate": "15-09-1997",
	"email": "patient@gmail.com",
	"password": "12345678"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register", reqBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
