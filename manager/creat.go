package creat

import (
	"log"
	"net/http"
)

type client struct {
	name   string
	number string
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
	}
}
