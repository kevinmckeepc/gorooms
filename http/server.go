package http

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"procore.com/rooms/ws"
)

type PageData struct {
	VisitorCount int
}

var (
	visitorCounter int
	tmpl = template.Must(template.ParseFiles("web/templates/index.html"))
	hub = ws.SharedHub()
)

func init() {
  initRoutes()
	// Run the hub in a new goroutine
	go hub.Run()
}

// Initialize the http routes
func initRoutes() {
  http.HandleFunc("/", indexRoute)
	http.HandleFunc("/data", dataRoute)
	http.HandleFunc("/ws", hub.Upgrade)
}

/// Returns the index.html
func indexRoute(w http.ResponseWriter, r *http.Request) {
	// Don't count the /favicon.ico request
	if r.URL.String() == "/" {
		visitorCounter++
	}

	tmpl.Execute(w, data()) // Merge the template & data
}

// Returns the raw visitor data in json
func dataRoute(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal(data())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Builds the page data to push into responses
func data() PageData {
	data := PageData {
		VisitorCount: visitorCounter,
	}
	return data
}

func Start() {
  fmt.Println("Starting http server ...")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
