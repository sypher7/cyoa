package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/sypher7/cyoa"
)

var tmpl *template.Template
var defaultTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
  <h1>{{.Title}}</h1>
  {{range .Paragraphs}}
    <p>{{.}}</p>
  {{end}}
  <ul>
  {{range .Options}}
    <li>
      <a href="/{{.Chapter}}">{{.Text}}</a>
    </li>
  {{end}}
  </ul>
</html>
`

type storyHandler struct {
	s *cyoa.Story
}

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	story := *h.s
	err := tmpl.Execute(w, story["intro"])
  if err != nil {
    log.Fatal(err)
  }
}

func NewStoryHandler(s *cyoa.Story) http.Handler {
	return storyHandler{s}
}

func init() {
	tmpl = template.Must(template.New("").Parse(defaultTemplate))
}

func main() {
	port := flag.Int("port", 3000, "the port for the web server")
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

	handler := NewStoryHandler(story)
  log.Printf("Starting the server on port: %d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), handler)
}
