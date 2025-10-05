package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	Id          int    `bson:"id"`
	OriginalURL string `bson:"OriginalURL"`
	ShortURL    string `bson:"ShortURL"`
}

//insert code

func InsertURL(ctx context.Context, collection *mongo.Collection, url *URL)  error {
	_, err := collection.InsertOne(ctx, url)
	return err
}

func FindURLbyShortCode(ctx context.Context, collection *mongo.Collection, shortCode string) (*URL, error){
	var url URL
	err := collection.FindOne(ctx, map[string]interface{}{"ShortURL": shortCode}).Decode(&url)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("not found")
	}

	return &url, nil
}