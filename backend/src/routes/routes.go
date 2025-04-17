package routes

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong!")
	})
}
