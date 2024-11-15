package left

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/JoseFredes/learn-go-hex/internal/core/ports"
)

type UserHandler struct {
	service ports.UserServices
}

func NewUserHandler(service ports.UserServices) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
    if idParam == "" {
        http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
        return
    }

	id, err := strconv.Atoi(idParam)
	if err != nil{
		http.Error(w, "Id is not a number", http.StatusBadRequest)
	}

	users, err := h.service.Get(id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
