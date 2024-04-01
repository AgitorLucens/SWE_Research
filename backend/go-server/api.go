package main

import (
	"log"
	"net/http"
)


type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("POST /users",func(w http.
		ResponseWriter, r *http.Request){
			userId :="fiorella hola"
			w.Write([]byte("User ID: " + userId))
		})

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}