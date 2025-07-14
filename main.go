package main

import (
	"log"
	"net/http"

	"encoding/json"

	address "github.com/EnnioSimoes/1-Deploy-CloudRun/address"
	weather "github.com/EnnioSimoes/1-Deploy-CloudRun/weather"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addr, err := address.GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid zipcode"})
		log.Println("Error getting address:", err)
		return
	}

	if addr.Cep == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "can not find zipcode"})
		log.Println("Error: Address not found for zipcode:", cep)
		return
	}

	temperature, err := weather.GetWeather(addr.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/temperature/{cep}", handler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		println("Ocorreu um erro: %w", err)
	}
}
