package promo

import (
	"GetfitWithPhysio-backend/helper"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func setupDabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/getfitwith_physio")
	helper.HandleError(err)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
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
