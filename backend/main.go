package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/LikheKeto/daraz-bazaar/scraper"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func match(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	title := r.URL.Query().Get("title")
	price := r.URL.Query().Get("price")
	category := r.URL.Query().Get("category")
	brand := r.URL.Query().Get("brand")
	if title == "" || price == "" {
		http.Error(w, "Invalid data!", http.StatusBadRequest)
		return
	}
	results, err := scraper.Scrape(scraper.TrimTitle(title), category, brand)
	if err != nil {
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/match", match)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
