package main

import (
	"encoding/json"
	"net/http"

	"github.com/anandxmj/weather-service/internal/weather/providers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	weatherService := &providers.NationalWeather{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/weather/{latitude},{longitude}", func(w http.ResponseWriter, r *http.Request) {
		latitude := chi.URLParam(r, "latitude")
		longitude := chi.URLParam(r, "longitude")

		weather, err := weatherService.GetWeather(latitude, longitude)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(weather); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}
