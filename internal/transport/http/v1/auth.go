package v1

import (
	"encoding/json"
	"net/http"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

func (t transport) Login(w http.ResponseWriter, r *http.Request) {
	var customer request.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jwtToken, err := t.s.Login(r.Context(), customer.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.Login{Access: jwtToken})
}
