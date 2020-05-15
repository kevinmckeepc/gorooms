package http

import (
	"fmt"
	"log"
  "net/http"
	"html/template"
	"encoding/json"
)

type PageData struct {
	VisitorCount int
}

var visitorCounter int
var tmpl = template.Must(template.ParseFiles("web/templates/index.html"))

func init() {
  initRoutes()
}

func initRoutes() {
  http.HandleFunc("/", indexRoute)
	http.HandleFunc("/data", dataRoute)
}

/// Returns the index.html
func indexRoute(w http.ResponseWriter, r *http.Request) {
	// Don't count the /favicon.ico request
	if r.URL.String() == "/" {
		visitorCounter++
	}

	data := PageData {
		VisitorCount: visitorCounter,
	}

	tmpl.Execute(w, data) // Merge the template & data
}

// Returns the raw visitor data in json
func dataRoute(w http.ResponseWriter, r *http.Request) {

	data := PageData {
		VisitorCount: visitorCounter,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Start() {
  fmt.Println("Starting http server")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
