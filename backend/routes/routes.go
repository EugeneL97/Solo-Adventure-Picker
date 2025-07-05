package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/models"
	"github.com/EugeneL97/solo-adventure-picker/services"
	"github.com/EugeneL97/solo-adventure-picker/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes() {

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		region := r.URL.Query().Get("region")
		col := config.Client.Database("solo-adventure-picker").Collection("adventures")

		pipeline := mongo.Pipeline{}

		if region != "" {
			pipeline = append(pipeline, bson.D{{"$match", bson.D{{"region", region}}}})
		}
		pipeline = append(pipeline, bson.D{{"$sample", bson.D{{"size", 1}}}})

		cursor, err := col.Aggregate(context.Background(), pipeline)
		if err != nil || !cursor.Next(context.Background()) {
			utils.WriteJSONError(w, http.StatusNotFound, utils.APIError{
				Error:   "No matching adventure found.",
				Code:    1001,
				Details: "Region either has no adventures or the database is down. Womp womp."})
			return
		}

		var adv models.Adventure
		if err := cursor.Decode(&adv); err != nil {
			utils.WriteJSONError(w, http.StatusInternalServerError, utils.APIError{
				Error:   "Error decoding adventure.",
				Code:    1002,
				Details: "You fed me improperly formatted JSON data. I hate you.",
			})
			return
		}

		// AI Enhancement: Create a basic user profile and enhance the adventure
		enhanceWithAI(&adv)

		json.NewEncoder(w).Encode(adv)
	})
}

// enhanceWithAI uses the adventure agent to personalize the adventure description
func enhanceWithAI(adv *models.Adventure) {
	ctx := context.Background()
	agent, err := services.NewAdventureAgent(ctx)
	if err != nil {
		log.Printf("Failed to create adventure agent: %v", err)
		// Set a default description if AI isn't available
		if adv.Description == "" {
			adv.Description = "Embark on this exciting adventure!"
		}
		// Set default XP value
		if adv.XPValue == 0 {
			adv.XPValue = 100
		}
		return
	}
	defer agent.Close()

	// Create a default user profile for now (later you'll get this from user data)
	userProfile := &services.UserProfile{
		Preferences: []string{"outdoor activities", "exploration", "mindfulness"},
		PastTypes:   []string{},
		EnergyLevel: "medium",
	}

	// Generate a basic description if none exists
	basicDescription := adv.Description
	if basicDescription == "" {
		basicDescription = "A wonderful adventure awaits you!"
	}

	// Enhance the description with AI
	enhanced, err := agent.EnhanceDescription(adv.Name, adv.Type, basicDescription, userProfile)
	if err != nil {
		log.Printf("Failed to enhance adventure description: %v", err)
		adv.Description = basicDescription
	} else {
		adv.Description = enhanced
	}

	// Set XP value based on adventure type (you can make this more sophisticated later)
	if adv.XPValue == 0 {
		switch adv.Type {
		case "hike", "outdoor":
			adv.XPValue = 150
		case "food", "cafe":
			adv.XPValue = 75
		case "culture", "museum":
			adv.XPValue = 125
		default:
			adv.XPValue = 100
		}
	}
}
