package models

//ShortenURL ...
type ShortenURL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
