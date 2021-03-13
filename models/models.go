package models

import (
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Pages     string             `json:"pages,omitempty" bson:"pages,omitempty"`
	Author    string             `json:"author" bson:"author"`
	Rating    int                `json:"rating" bson:"rating"`
	Reads     int                `json:"reads" bson:"reads"`
	Downloads int                `json:"downloads" bson:"reads"`
}

type Chapter struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookId primitive.ObjectID `json:"_book_id,omitempty" bson:"_book_id,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Page   int                `json:"page" bson:"page,omitempty"`
}

type Annexture struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookId  primitive.ObjectID `json:"_book_id,omitempty" bson:"_book_id,omitempty"`
	Content string             `json:"annexture" bson:"annexture,omitempty"`
}

type Error struct {
	Code int `json:"code"`
	Cause string `json:"cause"`
}