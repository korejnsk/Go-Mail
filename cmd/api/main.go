package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + " " + id))
		} else {
			w.Write([]byte("test"))
		}
	})
	r.Get("/{productName/{productId}}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(obj)
		w.Write(jsonData)
	})
	http.ListenAndServe(":3000", r)
}
