package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, ReadShowTemplate("hello"))
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ReadShowTemplate(resource string) string {
	html, err := os.ReadFile("templates/" + resource + "/show.html")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return string(html)
}
