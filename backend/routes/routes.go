package routes

import (
	"context"
	"encoding/json"
	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!"))
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")

		col := config.Client.Database("solo-adventure-picker").Collection("adventures")
		pipe := mongo.Pipeline{{{"$sample", bson.D{{"size", 1}}}}}

		cursor, err := col.Aggregate(context.Background(), pipe)
		if err != nil || !cursor.Next(context.Background()) {
			http.Error(w, "db error", 500)
			return
		}

		var adv models.Adventure
		if err := cursor.Decode(&adv); err != nil {
			http.Error(w, "decode error", 500)
			return
		}
		json.NewEncoder(w).Encode(adv)
	})
}
