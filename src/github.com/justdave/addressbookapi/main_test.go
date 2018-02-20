package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", create).Methods("GET")
	return router
}

func TestCreate(t *testing.T) {
	request, _ := http.NewRequest("GET", "/customer", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestRead(t *testing.T) {

}

func TestReadAll(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestDumpcsv(t *testing.T) {

}

func TestGrabcsv(t *testing.T) {

}

func TestParseobject(t *testing.T) {

}
