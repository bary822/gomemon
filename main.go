package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

func fetchMemoById(memoID MemoID) (Memo, error) {
	memoFixtures := []Memo{
		{
			ID:      1,
			Title:   "Grocery List",
			Content: "Milk, Bread, Eggs",
		},
		{
			ID:      2,
			Title:   "Meeting Notes",
			Content: "Reviewed project timeline and goals",
		},
		{
			ID:      3,
			Title:   "Vacation Ideas",
			Content: "Beach trip to Hawaii or ski trip to Aspen",
		},
	}

	for _, memo := range memoFixtures {
		if memo.ID == memoID {
			return memo, nil
		}
	}

	return Memo{}, fmt.Errorf("Memo with ID %d not found", memoID)
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	readTemplate("error", "404").Execute(w, nil)
}

func readTemplate(resource string, action string) *template.Template {
	t := template.Must(template.ParseFiles("templates/" + resource + "/" + action + ".html"))

	return t
}
