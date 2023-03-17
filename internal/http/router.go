package http

import "net/http"

func NewHandler() handler {
	var h handler
	return h
}

func (h *handler) Router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.RootHandle)
	mux.HandleFunc("/products", h.ProductsHandle)
}
