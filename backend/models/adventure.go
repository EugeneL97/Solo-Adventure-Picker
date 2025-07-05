package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Adventure struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"   json:"id"`
	Name        string             `bson:"name"            json:"name"`
	Type        string             `bson:"type"            json:"type,omitempty"`
	Region      string             `bson:"region"          json:"region"`
	Scenery     string             `bson:"scenery"         json:"scenery,omitempty"`
	Effort      string             `bson:"effort"          json:"effort,omitempty"`
	Duration    string             `bson:"duration"        json:"duration,omitempty"`
	Description string             `bson:"description"     json:"description,omitempty"`
	XPValue     int                `bson:"xpValue"         json:"xpValue,omitempty"`
}
