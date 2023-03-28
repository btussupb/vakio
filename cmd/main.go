package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/btussupb/vakio/internal/config"

	wayToHandlers "github.com/btussupb/vakio/internal/http"
	mngSrv "github.com/btussupb/vakio/internal/manager"
	wayToRepo "github.com/btussupb/vakio/internal/repository"
	wayToDb "github.com/btussupb/vakio/internal/storage"
)

func main() {
	fmt.Println("Server is listening in port localhost:8181")

	cfg, err := config.InitConfig("./config/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := wayToRepo.OpenDb(cfg.Database)
	if err != nil {
		fmt.Println(err, "main, opendb")
	}
	defer db.Close()

	storage := wayToDb.NewStorage(db)
	service := mngSrv.NewUserService(storage)

	rudiment := wayToHandlers.NewHandler(service)

	http.ListenAndServe(":8181", rudiment.Router())
}
