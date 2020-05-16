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

var (
	visitorCounter int
	tmpl = template.Must(template.ParseFiles("web/templates/index.html"))
	hub = newHub()
)

func init() {
  initRoutes()
	// Run the hub in a new goroutine
	go hub.run()
}

func initRoutes() {
  http.HandleFunc("/", indexRoute)
	http.HandleFunc("/data", dataRoute)
	http.HandleFunc("/ws", wsRoute)
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

// serveWs handles websocket requests from the peer.
func wsRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upgrading connection ...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("💩", err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
	// go client.echo()
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
