package http

import (
	"fmt"
	"log"
  "net/http"
	"html/template"
)

type PageData struct {
	VisitorCount int
}

var visitorCounter int

func init() {
  initRoutes()
}

func initRoutes() {

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Don't count the /favicon.ico request
		if r.URL.String() == "/" {
			visitorCounter++
		}

		data := PageData {
			VisitorCount: visitorCounter,
		}

		tmpl.Execute(w, data) // Merge the template & data
	})
}

func Start() {
  fmt.Println("Starting http server")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
