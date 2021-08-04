package utils

import (
	"encoding/json"
	"testing"

	"github.com/URLShortner/models"
)

func Test_checkFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"testfile", args{filename: "testdata.json"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("checkFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckIfExist(t *testing.T) {
	longURL := "https://www.testdata.com"
	shortURL := "http://localhost:8080/xVpzB5J"
	data := GetMockData()
	exist, existingshortUrl := CheckIfExist(data, longURL)
	if !exist || existingshortUrl.ShortURL != shortURL {
		t.Fatalf("it should exist")
	}
}
func TestCheckNotExist(t *testing.T) {
	longURL := "https://www.testdata1.com"
	shortURL := "http://localhost:8080/xVpzB5J"
	data := GetMockData()
	exist, existingshortUrl := CheckIfExist(data, longURL)
	if exist || existingshortUrl.ShortURL == shortURL {
		t.Fatalf("it should exist")
	}
}

//GetMockData ...
func GetMockData() (data []models.ShortenURL) {
	data = []models.ShortenURL{{LongURL: "https://www.testdata.com", ShortURL: "http://localhost:8080/xVpzB5J"}}
	err := WriteToFile(data, "test.json")
	if err != nil {
		return
	}
	return
}

func TestReadFile(t *testing.T) {
	GetMockData()
	filedata, err := ReadFile("test.json")
	if err != nil {
		t.Fatal("it should read the filedata")
	}
	data := []models.ShortenURL{}
	err = json.Unmarshal(filedata, &data)
	if err != nil {
		t.Fatal("it should read the data")
	}
}

func TestWriteTOFile(t *testing.T) {
	err := WriteToFile(GetMockData(), "test.json")
	if err != nil {
		t.Fatal("it should write the filedata")
	}
}
