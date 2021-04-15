package tests

import (
	"net/http"
	"net/http/httptest"
	"serverapp/webserver"
	"strings"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	file := "test.html"
	newFileHandler := &webserver.FilesHandler{FilePath: file}
	h := http.HandlerFunc(newFileHandler.HomePageHandler)
	h.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code retutned: %v", status)
	}
	expectedInBody := "Test Page"
	if body := rr.Body.String(); !strings.Contains(body, expectedInBody) {
		t.Error("Expected string is not present in body.")
	}
}
