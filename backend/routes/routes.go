package routes

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"

	"github.com/EugeneL97/solo-adventure-picker/models"
)

func RegisterRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		seeds := []models.Adventure{
			{Name: "El Corte de Madera Creek Preserve", Type: "hike"},
			{Name: "Mount Tamalpais", Type: "hike"},
			{Name: "Sunol Regional Wilderness", Type: "hike"},
			{Name: "Windy Hill Open Space Preserve", Type: "hike"},
			{Name: "Castle Rock State Park", Type: "hike"},
			{Name: "Purisima Creek Redwoods", Type: "hike"},
			{Name: "Lake Chabot Loop", Type: "hike"},
			{Name: "Edgewood Park and Natural Preserve", Type: "hike"},
			{Name: "Monte Bello Open Space Preserve", Type: "hike"},
			{Name: "Big Basin Redwoods", Type: "hike"},
		}

		random := seeds[rand.IntN(len(seeds))]
		json.NewEncoder(w).Encode(random)

	})
}
