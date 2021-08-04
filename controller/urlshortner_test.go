package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/URLShortner/models"
	"github.com/URLShortner/utils"
)

func TestMakeShorterURL(t *testing.T) {
	var jsonStr = []byte(`{"long_url":"https://www.test.com"}`)

	req, err := http.NewRequest("POST", "/api/url/short", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeShorterURL)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestMakeShorterURLAlredyExist(t *testing.T) {
	InsertMockData()
	var jsonStr = []byte(`{"long_url":"https://www.testdata.com"}`)
	req, err := http.NewRequest("POST", "/api/url/short", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeShorterURL)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//InsertMockData ...
func InsertMockData() {
	data := []models.ShortenURL{{LongURL: "https://www.testdata.com", ShortURL: "http://localhost:8080/xVpzB5J"}}
	utils.WriteToFile(data, "test.txt")
}

func TestGetAllShorterURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/url/short", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllURLs)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestMakeShorterURLErrorCondition(t *testing.T) {
	var jsonStr = []byte(`"testdata"`)

	req, err := http.NewRequest("POST", "/api/url/short", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeShorterURL)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestGetActualUrl(t *testing.T) {
	InsertMockData()
	req, err := http.NewRequest("GET", "/xVpzB5J", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ActualEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMovedPermanently)
	}
}
