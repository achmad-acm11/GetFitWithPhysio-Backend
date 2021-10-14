package team

import (
	"GetfitWithPhysio-backend/helper"
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
	return Config(db, validate, router)
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
