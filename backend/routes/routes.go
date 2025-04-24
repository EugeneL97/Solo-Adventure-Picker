package routes

import (
	"context"
	"encoding/json"
	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/models"
	"github.com/EugeneL97/solo-adventure-picker/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
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
		json.NewEncoder(w).Encode(adv)
	})
}
