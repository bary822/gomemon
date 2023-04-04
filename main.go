package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/memo", memoHandler)

	http.HandleFunc("/", notFoundHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func memoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, readTemplate("memo", "show"))
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, readTemplate("error", "404"))
}

func readTemplate(resource string, action string) string {
	html, err := os.ReadFile("templates/" + resource + "/" + action + ".html")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return string(html)
}
