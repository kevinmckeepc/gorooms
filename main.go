package main

import (
	"fmt"
	"procore.com/rooms/http"
	"procore.com/rooms/database"
)

func main() {
	fmt.Println("Running main")
	database.Connect()
	http.Start()

}
