package service

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

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/getfitwith_physio")

	helper.HandleError(err)

	return db
}

func setupRouter(db *sql.DB) *httprouter.Router {
	// Init validator
	validate := validator.New()

	// Init Router
	router := httprouter.New()

	return Config(db, validate, router)
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
