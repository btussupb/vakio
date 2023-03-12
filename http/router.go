package handlers

import "net/http"

func (h *handler) Router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.RootHandle)
}
