package main

import (
	"bitespeed-task/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Route("/api/v1/",func(r chi.Router) {
		serviceStore := services.NewStore()
		identityHandler:=services.NewHandler(serviceStore)
		identityHandler.RegisterRoutes(r)
	})

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080",router)

}