package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/memo", MemoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MemoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, ReadShowTemplate("memo"))
}

func ReadShowTemplate(resource string) string {
	html, err := os.ReadFile("templates/" + resource + "/show.html")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return string(html)
}
