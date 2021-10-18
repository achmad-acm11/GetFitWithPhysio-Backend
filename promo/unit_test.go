package promo

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

func setupDabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/getfitwith_physio?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	helper.HandleError(err)

	return db
}

func setupRouter(db *gorm.DB) http.Handler {
	// Init Validator
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	// Config End Point
	router = Config(db, validate, router)
	router.PanicHandler = exception.ErrorHandler
	return router
}
func truncatePromo(db *gorm.DB) {
	db.Exec("TRUNCATE promos")
}
func TestGetAllPromoSuccess(t *testing.T) {
	db := setupDabase()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/promos", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreatePromoSuccess(t *testing.T) {
	db := setupDabase()
	truncatePromo(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"discount":20}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/promos", reqBody)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	resBody, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(resBody, &responseBody)

	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
func TestCreatePromoFailed(t *testing.T) {
	db := setupDabase()
	truncatePromo(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"discount":0}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/promos", reqBody)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	resBody, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(resBody, &responseBody)

	assert.Equal(t, 400, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
