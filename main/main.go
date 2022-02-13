package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"dansdomain.net/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port number")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
