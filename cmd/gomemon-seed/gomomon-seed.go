package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:password@localhost:27017"

type MemoID int

type Memo struct {
	ID      MemoID
	Title   string
	Content string
}

func main() {
	seedMemos()
}

func seedMemos() {
	// set up MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// select the database and collection to use
	db := client.Database("mydb")
	memosCollection := db.Collection("memos")

	// insert some sample data
	memos := []interface{}{
		Memo{
			ID:      1,
			Title:   "Grocery List",
			Content: "Milk, Bread, Eggs",
		},
		Memo{
			ID:      2,
			Title:   "Meeting Notes",
			Content: "Reviewed project timeline and goals",
		},
		Memo{
			ID:      3,
			Title:   "Vacation Ideas",
			Content: "Beach trip to Hawaii or ski trip to Aspen",
		},
	}
	insertResult, err := memosCollection.InsertMany(context.Background(), memos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents\n", len(insertResult.InsertedIDs))
}
