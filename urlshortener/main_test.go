package main

import (
	"net/http"
	"testing"
)

func TestGetOriginalURL(t *testing.T) {
	// make a dummy request
	response, err := gttp.Get("http://localhost:8000/v1/short/1")

	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}
	if err != nil {
		r.Errorf("Encountered an error:", err)
	}
}
