package main

import (
	"fmt"
	"net/http"

	wayToHandlers "github.com/btussupb/vakio/internal/http"
	mngSrv "github.com/btussupb/vakio/internal/manager"
	// wayToDb "github.com/btussupb/vakio/db"
)

func main() {
	fmt.Println("Server is listening in port localhost:8181")

	db, err := OpenDb(config)
	service := mngSrv.NewUserService(db)
	rudiment := wayToHandlers.NewHandler(service)
	http.ListenAndServe(":8181", rudiment.Router())
}
