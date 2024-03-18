package v1

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"vk/internal/errs"
	"vk/internal/transport/http/response"
)

func parseQuery(urls url.Values, reqForm interface{}) error {

	data, err := json.Marshal(urls)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, reqForm); err != nil {
		return err
	}
	return nil
}
func handleError(err error) error {
	switch err.Error() {
	case errs.ServerError, errs.AuthError, errs.ValidationError, errs.InvalidToken, errs.NotAdmin, errs.NotLogged:
	default:
		err = fmt.Errorf(errs.UnexceptedError)

	}
	return err
}
func handleErrorAuth(err error) error {
	switch err.Error() {
	case errs.InvalidToken, errs.NotAdmin, errs.NotLogged:
	default:
		err = fmt.Errorf(errs.AuthError)

	}
	return err
}

func (t transport) adminPermission(w http.ResponseWriter, r *http.Request) error {
	jwtToken := r.Header.Get("Authorization")
	admin, err := t.s.IsAdmin(r.Context(), jwtToken)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}
	if !admin {
		//http.Error(w, , http.StatusUnauthorized)
		return fmt.Errorf(errs.NotAdmin)
	}
	return nil
}
func (t transport) userPermission(w http.ResponseWriter, r *http.Request) error {
	jwtToken := r.Header.Get("Authorization")
	res, err := t.s.IsLogged(r.Context(), jwtToken)

	if err != nil {
		return err
	}
	if !res {
		return fmt.Errorf(errs.NotLogged)
	}
	return nil
}

func (t transport) handleError(w http.ResponseWriter, err, handledError error, method string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response.Error{Error: handledError.Error()})
	t.l.Error(method, zap.Error(err))
}

//	func serverError(w http.ResponseWriter,  method string, status int){
//		w.WriteHeader(status)
//		json.NewEncoder(w).Encode(response.Error{Error: err.Error()})
//		t.l.Error(method, zap.Error(err))
//	}
func (t transport) handleOk(w http.ResponseWriter, resp interface{}, method string, status int) {

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
	t.l.Info(method)
}
