package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}
	if !isAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return err
	}
	return nil
}
