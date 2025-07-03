package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-service/services"
)

type RegisterRequest struct {
	Email    string `json: "email`
	Password string `json: "password"`
}

type LoginRequest struct {
	Email    string `json: "email`
	Password string `json: "password"`
}

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userId, err := h.UserService.RegisterUser(req.Email, req.Password)

	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"user_id": userId,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.UserService.LoginUser(req.Email, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
        return
	}

	response := map[string]interface{}{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) ProtectedResource(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the protected route!"))
}

