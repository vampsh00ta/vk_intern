package v1

import (
	"encoding/json"
	"net/http"
	"vk/internal/repository/models"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

func (t transport) AddActor(w http.ResponseWriter, r *http.Request) {
	if err := t.adminPermission(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actorReq request.AddActor

	if err := json.NewDecoder(r.Body).Decode(&actorReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	actor := models.Actor{
		Name:      actorReq.Name,
		Gender:    actorReq.Gender,
		BirthDate: actorReq.BirthDate,
		Films:     actorReq.Films,
	}
	id, err := t.s.AddActor(r.Context(), actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.AddActor{Id: id})
}
func (t transport) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if err := t.adminPermission(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var acrorReq request.DeleteActor

	if err := json.NewDecoder(r.Body).Decode(&acrorReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := t.s.DeleteActorByID(r.Context(), acrorReq.Id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(response.AddFilm{Id: id})
}
func (t transport) GetActors(w http.ResponseWriter, r *http.Request) {

	//var actorReq request.GetActors
	//
	//if err := json.NewDecoder(r.Body).Decode(&actorReq); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	actors, err := t.s.GetActors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.GetActors{Actors: actors})
}
