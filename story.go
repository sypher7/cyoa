package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var defaultHandlerTempl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
				<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    </head>

    <body>
        <div class="card" style="margin: 3em;">
				<div class="card-body">
						<h5 class="card-title">{{.Title}}</h5>
						{{range .Story}}
								<p class="card-text">{{.}}</p>
						{{end}}

						{{range .Options}}
								<a href="/{{.Chapter}}" class="card-link">{{.Text}}</a><br />
						{{end}}
				</div>
		</div>
    </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		tpl := template.Must(template.New("").Parse(defaultHandlerTempl))
		err := tpl.Execute(w, chapter)

		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func JsonStory(r io.Reader) (*Story, error) {
	var story *Story

	d := json.NewDecoder(r)

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
