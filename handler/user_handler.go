package handler

import (
	"demo/dto"
	"demo/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	svc service.AdminService
}

func NewAdminHandler(svc service.AdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}

func (h *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	res, err := h.svc.CreateAdmin(req)
	if err != nil {
		http.Error(w, "failed to create admin", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, err := h.svc.GetAdmin(id)
	if err != nil {
		http.Error(w, "admin not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(res)
}
