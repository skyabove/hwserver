package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

const remoteIpKey = "remote_ip"

func routes() http.Handler {
	mux := chi.NewRouter()

	//use middleware
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	//routing
	mux.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello man, your IP is:" + request.RemoteAddr))
	})

	return mux
}
