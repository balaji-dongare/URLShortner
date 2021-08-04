package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/URLShortner/models"
	"github.com/sirupsen/logrus"
)

const FilenName = "test.json"

//ReadFile read the file data
func ReadFile(filename string) (filedata []byte, err error) {
	err = checkFile(filename)
	if err != nil {
		return
	}
	filedata, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return
}

//WriteToFile write the URL data into file
func WriteToFile(data []models.ShortenURL, filename string) (err error) {
	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}
	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		logrus.Error(err)
	}
	return
}

//CheckIfExist check the short url is alredy exist for input long url
func CheckIfExist(data []models.ShortenURL, long_url string) (exist bool, url models.ShortenURL) {
	for _, urls := range data {
		if urls.LongURL == long_url {
			exist = true
			url = urls
			return
		}
	}
	return
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
