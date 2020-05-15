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
		visitorCounter++
			data := PageData {
				VisitorCount: visitorCounter,
			}
			tmpl.Execute(w, data)
	})
}

func Start() {
  fmt.Println("Starting http server")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
