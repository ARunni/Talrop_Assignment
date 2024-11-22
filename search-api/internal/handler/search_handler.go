package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search-api/internal/domain"
	"search-api/internal/usecase/interfaces"
	"strconv"
)

type SearchHandler struct {
	useCase interfaces.Usecase
}

func NewSearchHandler(uc interfaces.Usecase) *SearchHandler {
	return &SearchHandler{useCase: uc}
}

func (h *SearchHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println("name ==", name)
	if name == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	pagestr := r.URL.Query().Get("page")
	// offsetStr := r.URL.Query().Get("offset")

	page, err := strconv.Atoi(pagestr)

	if err != nil {
		http.Error(w, "Missing 'page' query parameter", http.StatusBadRequest)
		return
	}

	limit := 100

	offset := (page - 1) * limit

	users, err := h.useCase.SearchUsersByName(name, limit, offset)
	if err != nil {
		http.Error(w, "Error retrieving users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Total int           `json:"total"`
		Users []domain.User `json:"users"`
	}{
		Total: len(users),
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
