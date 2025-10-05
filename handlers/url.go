package handlers

import (
	"UrlShortner/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"

	"go.mongodb.org/mongo-driver/mongo"
)

var URLCollection *mongo.Collection

func GenerateShortCode() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const length = 6
    
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func ShortenURLhandler(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		URL string `json: "URL"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil || reqData.URL == "" {
		http.Error(w, "invalid input or json value missing", http.StatusBadRequest)
		return
	}
	fmt.Println(reqData)

	shortCode := GenerateShortCode()

	url := models.URL{
		Id:          int(time.Now().Unix()),
		ShortURL:    shortCode,
		OriginalURL: reqData.URL,
	}

	err2 := models.InsertURL(r.Context(), URLCollection, &url)
	if err2 != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"shortcode": shortCode,
		"URL": "localhost:8080/redirect/" + shortCode,
		"OriginalUrl": reqData.URL,
	})
	

}

func Redirecthandler(w http.ResponseWriter, r *http.Request) {

	ShortCode := strings.TrimPrefix(r.URL.Path, "/redirect/")
	if ShortCode == "" {
		http.NotFound(w, r)
            return
	}
	fmt.Println(ShortCode)

	RedirectURL, err := models.FindURLbyShortCode(r.Context(), URLCollection, ShortCode)
	if err != nil {
		http.Error(w, "redirecturl error", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	Redirectedurl := string(RedirectURL.OriginalURL)
	if !strings.HasPrefix(Redirectedurl, "http://") && !strings.HasPrefix(Redirectedurl, "https://") {
		Redirectedurl = "https://" + Redirectedurl
	}
	http.Redirect(w, r, Redirectedurl, http.StatusFound)

}

func GetUrLDetailsHandler(w http.ResponseWriter, r *http.Request){
	prefix := "/url/"
	path := r.URL.Path

	if path == "" {
		http.NotFound(w,r)
		return
	}
	
	if !strings.HasPrefix(path, prefix) {
		http.NotFound(w,r)
		return
	}
	
	ShortCode := strings.TrimPrefix(path, prefix)
	
	if ShortCode == "" {
		http.Error(w, "enter a shortcode", http.StatusBadRequest)
		return
	}

	urlRecord, err := models.FindURLbyShortCode(r.Context(), URLCollection, ShortCode)
	if err != nil {
		http.Error(w,"error fetching your shorturl", http.StatusInternalServerError)
		return
	}

	fmt.Println(urlRecord)

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"ShortUrl" : "localhost:8080/redirect/" + urlRecord.ShortURL,
		"OriginalUrl": urlRecord.OriginalURL,
	})
}
