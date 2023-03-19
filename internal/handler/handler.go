package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"smartway-test-task/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/employee/create", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/employee/get", h.GetAllByCompany).Methods(http.MethodGet).Queries("company", "{id}")
	r.HandleFunc("/employee/get", h.GetAllByDepartment).Methods(http.MethodGet).Queries("department", "{name}")
	r.HandleFunc("/employee/update/{id}", h.Update).Methods(http.MethodPost)
	r.HandleFunc("/employee/delete/{id}", h.Delete).Methods(http.MethodPost)

	return r
}
