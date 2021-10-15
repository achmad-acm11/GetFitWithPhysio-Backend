package team

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

func setupRouter(db *gorm.DB) http.Handler {
	// Init Validate
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	// Congig Router and Create Endpoint
	router = Config(db, validate, router)

	router.PanicHandler = exception.ErrorHandler

	return router
}
func truncateTeam(db *gorm.DB) {
	db.Exec("TRUNCATE teams")
}
func TestGetAllTeamSuccess(t *testing.T) {
	db := setupDatabase()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/teams", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
func TestCreateTeamSuccess(t *testing.T) {
	db := setupDatabase()
	truncateTeam(db)
	router := setupRouter(db)

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	writer.WriteField("name", "Rifa Rahmalia. Amd.Ft")
	writer.WriteField("position", "therapist")
	writer.WriteField("url", "https://id.linkedin.com/in/faizah-abdullah-24a332b3")
	writer.WriteField("description", "Education And Training Manager di Lembaga Vokasi Universitas Indonesia.  Lecturer University of Indonesia.  Senior Physiotherapist OM Active Clinic")
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/teams", reqBody)
	request.Header.Add("content-type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
func TestCreateTeamFailed(t *testing.T) {
	db := setupDatabase()
	truncateTeam(db)
	router := setupRouter(db)

	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	writer.WriteField("name", "")
	writer.WriteField("position", "therapist")
	writer.WriteField("url", "https://id.linkedin.com/in/faizah-abdullah-24a332b3")
	writer.WriteField("description", "Education And Training Manager di Lembaga Vokasi Universitas Indonesia.  Lecturer University of Indonesia.  Senior Physiotherapist OM Active Clinic")
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/teams", reqBody)
	request.Header.Add("content-type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["meta"].(map[string]interface{})["code"].(float64)))
}
