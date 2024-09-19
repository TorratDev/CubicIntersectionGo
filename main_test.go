package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a request to pass to our handler. The second parameter is the URL path,
	// and the third parameter is the body (nil for GET requests).
	reqBody := `{
    "first": {
        "dimensions": {
            "x": 2,
            "y": 2,
            "z": 2
        },
        "center": {
            "x": 0,
            "y": 0,
            "z": 0
        }
    },
    "second": {
        "dimensions": {
            "x": 2,
            "y": 2,
            "z": 2
        },
        "center": {
            "x": 0.5,
            "y": 0.5,
            "z": 0.5
        }
    }
}`
	req, err := http.NewRequest("POST", "/intersect", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder, which will capture the response from the handler.
	rr := httptest.NewRecorder()

	intersect := intersectServiceImpl{}

	// Create a handler function and pass it to ServeHTTP, simulating an HTTP request.
	handler := cubicHandler(intersect)
	handler.ServeHTTP(rr, req)

	// Check if the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := CubicResponse{}
	err = json.Unmarshal([]byte(`{"success":true,"volume":3.375}`), &expected)
	response := CubicResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)

	if response != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", response, expected)
	}
}
