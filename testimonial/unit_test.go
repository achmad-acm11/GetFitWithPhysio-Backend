package testimonial

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
	// Init Validator
	validate := validator.New()
	// Init Router
	router := httprouter.New()

	router = Config(db, validate, router)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func truncateTestimonial(db *gorm.DB) {
	db.Exec("TRUNCATE testimonials")
}

func TestGetAllTestimonialSuccess(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/testimonials", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateTestimonialSuccess(t *testing.T) {
	db := setupDatabase()
	truncateTestimonial(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"content": "Terimakasih"
	  }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/testimonials", reqBody)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
func TestCreateTestimonialFailed(t *testing.T) {
	db := setupDatabase()
	truncateTestimonial(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"content": ""
	  }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/testimonials", reqBody)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
