package http

import (
	"encoding/json"
	"html/template"
	"net/http"

	mngSrv "github.com/btussupb/vakio/internal/manager"
)

type handler struct {
	mngSrv mngSrv.Service
}

var temp *template.Template

func (h *handler) RootHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "template.html", nil)
	} else if r.Method == http.MethodPost {
		h.methodPost(w, r)
	}
}

func (h *handler) methodPost(w http.ResponseWriter, r *http.Request) {
	var userInput mngSrv.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		//
	}
	defer r.Body.Close()

	if err := h.mngSrv.PostUser(userInput); err != nil {
		//
	}

	w.Write([]byte("мы вам напишем"))
}
