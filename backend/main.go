package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/EugeneL97/solo-adventure-picker/config"
	"github.com/EugeneL97/solo-adventure-picker/routes"
)

func main() {
	config.InitDB(os.Getenv("MONGO_URI"))
	routes.RegisterRoutes()

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
