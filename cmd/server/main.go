package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"private-chat/internal/server/handler"
)

func main() {
	messageHandler := handler.NewMessageHandler()
	router := chi.NewRouter()
	router.Handle("/status", http.HandlerFunc(checkHealth))
	router.Route("/message", func(router chi.Router) {
		router.Post("/", http.HandlerFunc(messageHandler.HandleMessage))
	})
	http.ListenAndServe("localhost:8080", router)
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint("I am good! And what about you")))
}
