package v1

import (
	"encoding/json"
	"github.com/gorilla/schema"
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
func (t transport) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	var filmReq request.DeleteFilm

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

	if err := t.s.DeleteFilm(r.Context(), filmReq.Id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(response.AddFilm{Id: id})
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
	var queryForm request.GetFilm
	if err := schema.NewDecoder().Decode(&queryForm, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	queryForm.SortBy = r.URL.Query().Get("sort_by")
	queryForm.OrderBy = r.URL.Query().Get("order_by")
	queryForm.Name = r.URL.Query().Get("name")
	queryForm.Title = r.URL.Query().Get("title")

	var films []models.Film
	if queryForm.Title != "" {
		films, err = t.s.GetFilmsByTitle(r.Context(), queryForm.Title, queryForm.SortBy, queryForm.OrderBy)
	} else if queryForm.Name != "" {
		films, err = t.s.GetFilmsByActorName(r.Context(), queryForm.Name, queryForm.SortBy, queryForm.OrderBy)
	} else {
		films, err = t.s.GetFilms(r.Context(), queryForm.SortBy, queryForm.OrderBy)

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.GetFilms{Films: films})
}
