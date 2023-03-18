package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	mngSrv "github.com/btussupb/vakio/internal/manager"
)

type handler struct {
	mngSrv mngSrv.Service
}

var temp *template.Template

func (h *handler) RootHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		temp, err := template.ParseFiles("front/index.html")
		if err != nil {
			fmt.Println("ParseFiles")
		}
		err = temp.ExecuteTemplate(w, "index", nil)
		if err != nil {
			fmt.Println("ExecuteTemplate")
		}
	case http.MethodPost:
		h.methodPost(w, r)
	}
}

func (h *handler) methodPost(w http.ResponseWriter, r *http.Request) {
	var userInput mngSrv.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		fmt.Println("Decode r.Body")
	}
	defer r.Body.Close()

	if err := h.mngSrv.PostUser(userInput); err != nil {
		w.Write([]byte("ошибка на сервере"))
	} else {
		w.Write([]byte("мы вам напишем"))
	}
}
