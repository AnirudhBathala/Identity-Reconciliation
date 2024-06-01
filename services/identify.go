package services

import (
	"bitespeed-task/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store models.ServiceStore
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Post("/identify",h.HandleIdentify)
}

func (h *Handler) HandleIdentify(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Request sucessfull"))
}

func NewHandler(s models.ServiceStore) *Handler {
	return &Handler{
		store: s,
	}
}
