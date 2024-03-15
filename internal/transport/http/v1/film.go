package v1

import (
	"encoding/json"
	"net/http"
	"vk/internal/repository/models"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

//func GetFilms(w http.ResponseWriter, r *http.Request) {
//	var p Person
//
//	// Try to decode the request body into the struct. If there is an error,
//	// respond to the client with the error message and a 400 status code.
//	err := json.NewDecoder(r.Body).Decode(&p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// Do something with the Person struct...
//	fmt.Fprintf(w, "Person: %+v", p)
//}

func (t transport) AddFilm(w http.ResponseWriter, r *http.Request) {
	var filmReq request.AddFilm

	jwtToken := r.Header.Get("Authorization")
	isAdmin, err := t.s.IsAdmin(r.Context(), jwtToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !isAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&filmReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	film := models.Film{
		Actors:      filmReq.Actors,
		Title:       filmReq.Title,
		Rating:      filmReq.Rating,
		Description: filmReq.Description,
		ReleaseDate: filmReq.ReleaseDate,
	}
	id, err := t.s.AddFilm(r.Context(), film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.AddFilm{Id: id})
}

func (t transport) GetFilms(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	isAdmin, err := t.s.IsAdmin(r.Context(), jwtToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !isAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return
	}

	films, err := t.s.GetFilms(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.GetFilms{Films: films})
}