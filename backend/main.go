package main

import (
	"fmt"
	"github.com/EugeneL97/solo-adventure-picker/routes"
	"net/http"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
