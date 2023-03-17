package main

import (
	"fmt"
	"net/http"

	wayToHandlers "github.com/btussupb/vakio/internal/http"
	// wayToDb "github.com/btussupb/vakio/db"
)

func main() {
	fmt.Println("Server is listening in port localhost:8181")
	rudiment := wayToHandlers.NewHandler()
	http.ListenAndServe(":8181", rudiment.Router())
}
