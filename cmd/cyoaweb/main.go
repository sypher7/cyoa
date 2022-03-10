package main

import (
	"flag"
	"log"
	"os"

  "github.com/sypher7/cyoa"
)

func main()  {
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

  log.Printf("%+v\n", *story)
}
