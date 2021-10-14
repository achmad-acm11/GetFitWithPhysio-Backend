package user

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

	// Config End Point
	router = Config(db, validate, router)

	// Handler Error
	router.PanicHandler = exception.ErrorHandler

	return router
}

func TestLoginSuccess(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
	"email": "patient@gmail.com",
    "password": "12345678"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", reqBody)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}

func TestLoginFailed(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"email": "patient1@gmail.com",
		"password": "12345678"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", reqBody)
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
