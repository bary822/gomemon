package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/memo", MemoHandler)

	http.HandleFunc("/", NotFoundHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MemoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, ReadTemplate("memo", "show"))
}

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, ReadTemplate("error", "404"))
}

func ReadTemplate(resource string, action string) string {
	html, err := os.ReadFile("templates/" + resource + "/" + action + ".html")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return string(html)
}
