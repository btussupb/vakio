package main

import (
	"fmt"
	"net/http"

	wayToHandlers "github.com/btussupb/vakio/handlers"
)

func main() {
	fmt.Println("Server is listening in port 8181")
	wayToHandlers.Router()
	http.ListenAndServe(":8181", nil)
}
