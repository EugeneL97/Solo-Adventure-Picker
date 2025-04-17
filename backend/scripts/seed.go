package main

import (
	"context"
	"log"

	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/models"
)

func main() {
	config.InitDB("mongodb://localhost:27017")
	col := config.Client.Database("solo-adventure-picker").Collection("adventures")

	seed := []interface{}{
		models.Adventure{Name: "El Corte de Madera Creek Preserve", Type: "hike"},
		models.Adventure{Name: "Mount Tamalpais", Type: "hike"},
		models.Adventure{Name: "Sunol Regional Wilderness", Type: "hike"},
		models.Adventure{Name: "Windy Hill Open Space Preserve", Type: "hike"},
		models.Adventure{Name: "Castle Rock State Park", Type: "hike"},
		models.Adventure{Name: "Purisima Creek Redwoods", Type: "hike"},
		models.Adventure{Name: "Lake Chabot Loop", Type: "hike"},
		models.Adventure{Name: "Edgewood Park and Natural Preserve", Type: "hike"},
		models.Adventure{Name: "Monte Bello Open Space Preserve", Type: "hike"},
		models.Adventure{Name: "Big Basin Redwoods", Type: "hike"},
	}
	if _, err := col.InsertMany(context.Background(), seed); err != nil {
		log.Fatal(err)
	}
	log.Println("Seed inserted!")
}
