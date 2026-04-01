package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	router := SetupRouter()

	body := []byte(`{"username": "Admin", "age": 10}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	
	bodyOk := []byte(`{"username": "Admin", "age": 25}`)
	wOk := httptest.NewRecorder()
	reqOk, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyOk))
	router.ServeHTTP(wOk, reqOk)
	assert.Equal(t, 200, wOk.Code)
}