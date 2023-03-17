package main

import (
	"fmt"
	"net/http"

	wayToHandlers "github.com/btussupb/vakio/internal/http"
)

func main() {
	fmt.Println("Server is listening in port localhost:8181")
	rudiment := wayToHandlers.NewHandler()
	rudiment.Router()
	http.ListenAndServe(":8181", nil)
}
