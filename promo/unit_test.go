package promo

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
	return Config(db, validate, router)
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
