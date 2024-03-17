package v1

import (
	"encoding/json"
	"net/http"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

func (t transport) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "Login"
	var customer request.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)

		return
	}
	jwtToken, err := t.s.Login(r.Context(), customer.Username)
	if err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.Login{Access: jwtToken}, methodName, http.StatusOK)

}
