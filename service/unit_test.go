package service

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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
	// Init validator
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	router = Config(db, validate, router)

	router.PanicHandler = exception.ErrorHandler

	return router
}
func truncateService(db *gorm.DB) {
	db.Exec("TRUNCATE services")
}
func TestGetAllServiceSuccess(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/services", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateServiceSuccess(t *testing.T) {
	db := setupDatabase()
	truncateService(db)
	router := setupRouter(db)

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	writer.WriteField("kode_promo", "")
	writer.WriteField("service_name", "Product Gold")
	writer.WriteField("kuota_meet", "2")
	writer.WriteField("price", "450000")
	writer.WriteField("description", "Treatment Fisioterapi reguler dengan 1x pertemuan.")
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/services", reqBody)
	request.Header.Add("content-type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}

func TestCreateServiceFailed(t *testing.T) {
	db := setupDatabase()
	truncateService(db)
	router := setupRouter(db)

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	writer.WriteField("kode_promo", "")
	writer.WriteField("service_name", "")
	writer.WriteField("kuota_meet", "2")
	writer.WriteField("price", "450000")
	writer.WriteField("description", "Treatment Fisioterapi reguler dengan 1x pertemuan.")
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/services", reqBody)
	request.Header.Add("content-type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
