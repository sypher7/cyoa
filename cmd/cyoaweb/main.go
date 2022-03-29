package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sypher7/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web app on")
	filename := flag.String("file", "./data/story.json", "the JSON file with the CYOA story")

	log.Printf("Using the story file: %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Unable to open story file")
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatal("Unable to decode story file")
	}

	h := cyoa.NewHandler(*story, nil)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
