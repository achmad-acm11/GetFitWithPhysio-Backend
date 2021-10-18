package transaction

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

func truncateTransaction(db *gorm.DB) {
	db.Exec("TRUNCATE transactions")
}

func TestCreateTransactionSucces(t *testing.T) {
	db := setupDatabase()
	truncateTransaction(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"describe_complaint": "Sakit pada bagian pinggang"
	  }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/transactions/1", reqBody)
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
func TestCreateTransactionFailed(t *testing.T) {
	db := setupDatabase()
	truncateTransaction(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"describe_complaint": ""
	  }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/transactions/1", reqBody)
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
func TestCreateTransactionNotFound(t *testing.T) {
	db := setupDatabase()
	truncateTransaction(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"describe_complaint": "Sakit pada bagian pinggang"
	  }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/transactions/999", reqBody)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
