package left

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JoseFredes/learn-go-hex/internal/infra/left/dto"
	service "github.com/JoseFredes/learn-go-hex/internal/services/user"
)

type UserHandler struct {
	service service.UserServices
}

func NewUserHandler(service service.UserServices) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Id is not a number", http.StatusBadRequest)
	}

	users, err := h.service.Get(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse user", http.StatusBadRequest)
		return
	}

	if err := user.ValitateNewUser(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid user: %v", err), http.StatusBadRequest)
		return
	}

	userCreated, err := h.service.Create(user.ParseToDomain())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userCreated)
}