package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	server := NewServer("8080")

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server.healthHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("expected status 'healthy', got %s", response.Status)
	}

	if response.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got %s", response.Version)
	}
}

func TestHelloHandler(t *testing.T) {
	server := NewServer("8080")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server.helloHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response HelloResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if response.Message != "Hello, World!" {
		t.Errorf("expected message 'Hello, World!', got %s", response.Message)
	}
}

func TestNewServer(t *testing.T) {
	server := NewServer("8080")

	if server == nil {
		t.Error("NewServer returned nil")
		return
	}

	if server.server.Addr != ":8080" {
		t.Errorf("expected server address ':8080', got %s", server.server.Addr)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	server := NewServer("8080")

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Test the middleware by calling it directly
	server.loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("test")); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	})).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("middleware returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
