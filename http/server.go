package http

import (
	"fmt"
	"log"
  "net/http"
)

func init() {
  initRoutes()
}

func initRoutes() {
  http.HandleFunc("/", indexRoute)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, Hello())
}

// Hello ... The ever classic "Hello World"
func Hello() string {
    return "Hello, world"
}

func Start() {
  fmt.Println("Starting http server")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
