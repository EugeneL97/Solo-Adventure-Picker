package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"

	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/models"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	config.InitDB(os.Getenv("MONGO_URI"))
	col := config.Client.Database("solo-adventure-picker").Collection("adventures")

	n, err := col.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Println("Seed count error: ", err)
		return
	}

	if n > 0 {
		log.Printf("Adventures already seeded: %d\n", n)
		return
	}

	seed := []interface{}{
		models.Adventure{Name: "El Corte de Madera Creek Preserve", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Mount Tamalpais", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Sunol Regional Wilderness", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Windy Hill Open Space Preserve", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Castle Rock State Park", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Purisima Creek Redwoods", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Lake Chabot Loop", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Edgewood Park and Natural Preserve", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Monte Bello Open Space Preserve", Type: "hike", Region: "bay-area"},
		models.Adventure{Name: "Big Basin Redwoods", Type: "hike", Region: "bay-area"},
	}
	if _, err := col.InsertMany(context.Background(), seed); err != nil {
		log.Fatal(err)
	}
	log.Printf("Seed inserted with %d adventures!\n", len(seed))
}
