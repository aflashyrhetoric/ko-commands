package router

import (
	"log"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{mux: http.NewServeMux()}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Serve(addr string) {
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r.mux))
}
