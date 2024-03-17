package v1

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"vk/internal/transport/http/response"
)

func parseQuery(urls url.Values, reqForm interface{}) error {

	data, err := json.Marshal(urls)
	if err != nil {
		return err
	}
	fmt.Println(urls)
	if err = json.Unmarshal(data, reqForm); err != nil {
		return err
	}
	return nil
}
func (t transport) adminPermission(w http.ResponseWriter, r *http.Request) error {
	jwtToken := r.Header.Get("Authorization")
	isAdmin, err := t.s.IsAdmin(r.Context(), jwtToken)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}
	if !isAdmin {
		//http.Error(w, , http.StatusUnauthorized)
		return fmt.Errorf("not admin")
	}
	return nil
}

func (t transport) handleError(w http.ResponseWriter, err error, method string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response.Error{Error: err.Error()})
	t.l.Error(method, zap.Error(err))
}

func (t transport) handleOk(w http.ResponseWriter, resp interface{}, method string, status int) {

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
	t.l.Info(method)
}
