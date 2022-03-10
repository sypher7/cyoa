package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

  "github.com/sypher7/cyoa"
)

func main()  {
  filename := flag.String("file", "../../data/story1.json", "the JSON file with the CYOA story")

  fmt.Printf("Using the story file: %s\n", *filename)

  f, err := os.Open(*filename)
  if err != nil {
    log.Fatal("Unable to open story file")
  }

  d := json.NewDecoder(f)
  var story cyoa.Story
  if err := d.Decode(&story); err != nil {
    log.Fatal("Unable to decode story file")
  }

  log.Printf("%+v\n", story)
}
