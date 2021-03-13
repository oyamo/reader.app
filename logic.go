package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"reader.app/models"
	"time"
)

var (
	database = Client.Database("Reader")
)

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	// Perform Database
	collection := database.Collection("Book")
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	cursor, queryError := collection.Find(ctx, bson.M{})

	// If there is a query err, tell the user
	if queryError != nil {
		response := models.Error{
			Code:  500,
			Cause: "Internal server Error Occurred",
		}
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	if cursor != nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var book models.Book
			_ = cursor.Decode(&book)
			books = append(books, book)
		}
	}

	_ = json.NewEncoder(w).Encode(books)
}
