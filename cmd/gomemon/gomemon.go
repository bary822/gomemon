package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:password@localhost:27017"

func main() {
	http.HandleFunc("/memo/", memoHandler)

	http.HandleFunc("/", notFoundHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func memoHandler(w http.ResponseWriter, req *http.Request) {
	memoID, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	memo, err := fetchMemoById(MemoID(memoID))
	if err != nil {
		log.Println(err)
		notFoundHandler(w, req)
		return
	}

	readTemplate("memo", "show").ExecuteTemplate(w, "show.html", struct {
		Title   string
		Content string
	}{
		Title:   memo.Title,
		Content: memo.Content,
	})
}

type MemoID int

type Memo struct {
	ID      MemoID
	Title   string
	Content string
}

// fetch memo by id from mongodb
// if memo not found, return error
// if memo found, return memo
func fetchMemoById(memoID MemoID) (Memo, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("mydb")
	memosCollection := db.Collection("memos")
	filter := bson.D{{"id", memoID}}

	var memo Memo
	err = memosCollection.FindOne(context.Background(), filter).Decode(&memo)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return Memo{}, fmt.Errorf("Memo with ID %d not found", memoID)
		}
		log.Fatal(err)
		return Memo{}, fmt.Errorf("Internal server error")
	}

	return memo, nil
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	readTemplate("error", "404").Execute(w, nil)
}

func readTemplate(resource string, action string) *template.Template {
	t := template.Must(template.ParseFiles("../../web/template/" + resource + "/" + action + ".html"))

	return t
}
