package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"UrlShortner/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/rs/cors"
)

// var client *mongo.Client
var urlCollection *mongo.Collection
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("error connecting to database,", err)
		return
	}

	fmt.Println("db connected successfully")
	
	urlCollection = client.Database("urls").Collection("urls")
	//giving collection to handler file (it'll be assigned to URLCollection variable)
	handlers.URLCollection = urlCollection

	defer client.Disconnect(ctx)

	//router setup
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten/", handlers.ShortenURLhandler)
	mux.HandleFunc("/redirect/", handlers.Redirecthandler)
	mux.HandleFunc("/url/", handlers.GetUrLDetailsHandler)

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"}, // Next.js dev server
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })

	handler := c.Handler(mux)

	//start the server
	fmt.Println("Starting server on :8080")
    err = http.ListenAndServe(":8080", handler)
    if err != nil {
        fmt.Println(err)
    }

}