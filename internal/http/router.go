package http

import (
	"net/http"

	mngSrv "github.com/btussupb/vakio/internal/manager"
)

type service struct {
	mngSrv mngSrv.Service
}

func NewMngService(mngSrv mngSrv.Service) *service {
	return &service{mngSrv: mngSrv}
}

// func NewHandler(mngSrv mngSrv.Service) *handler {
// 	return &handler{
// 		mngSrv: mngSrv,
// 	}
// }

func (h *handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("./front/"))))
	// mux.Handle("/front/img/", http.StripPrefix("/front/img/", http.FileServer(http.Dir("./front/img/"))))

	mux.HandleFunc("/", h.RootHandle)
	mux.HandleFunc("/products", h.ProductsHandle)
	mux.Handle("/front/css/", http.StripPrefix("/front/css/", http.FileServer(http.Dir("./front/css"))))
	return mux
}
