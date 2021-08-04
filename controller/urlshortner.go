package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/URLShortner/models"
	"github.com/URLShortner/utils"
	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
)

//MakeShorterURL create a new short url if it dose not exists
func MakeShorterURL(w http.ResponseWriter, r *http.Request) {
	url := models.ShortenURL{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	filedata, err := utils.ReadFile(utils.FilenName)
	if err != nil {
		log.Println("not able to read a file")
	}
	data := []models.ShortenURL{}
	err = json.Unmarshal(filedata, &data)
	if err != nil {
		log.Println("empty file")
	}
	exist, existingurl := utils.CheckIfExist(data, url.LongURL)
	if exist {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(existingurl.ShortURL)
		return
	}
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	now := time.Now()
	hash, err := h.Encode([]int{int(now.Unix())})
	if err != nil {
		log.Println("not able to encode")
	}
	url.ShortURL = "http://localhost:8080/" + hash

	data = append(data, url)
	utils.WriteToFile(data, utils.FilenName)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(url.ShortURL)
}

//GetAllURLs Get all Urls
func GetAllURLs(w http.ResponseWriter, r *http.Request) {
	filedata, err := utils.ReadFile(utils.FilenName)
	if err != nil {
		log.Println("not able to read a file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := []models.ShortenURL{}
	json.Unmarshal(filedata, &data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

//ActualEndpoint redirect to actual url
func ActualEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	filedata, err := utils.ReadFile(utils.FilenName)
	if err != nil {
		log.Println("not able to read a file")
	}
	data := []models.ShortenURL{}
	err = json.Unmarshal(filedata, &data)
	if err != nil {
		log.Println("empty file")
	}
	url := utils.GetActualURLt(data, hash)
	log.Println(url)
	if url != "" {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
