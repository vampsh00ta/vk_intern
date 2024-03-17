package v1

import (
	"encoding/json"
	"net/http"
	"vk/internal/repository/models"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

func (t transport) AddActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "AddActor"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)
		return
	}

	var actorReq request.AddActor

	if err := json.NewDecoder(r.Body).Decode(&actorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)
		return
	}
	actor := models.Actor{
		Name:      actorReq.Name,
		Gender:    actorReq.Gender,
		BirthDate: actorReq.BirthDate,
		FilmIds:   actorReq.Films,
	}
	id, err := t.s.AddActor(r.Context(), actor)
	if err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)
		return
	}

	t.handleOk(w, response.AddActor{Id: id}, methodName, http.StatusCreated)

}
func (t transport) DeleteActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "DeleteActor"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)
		return
	}

	var acrorReq request.DeleteActor

	if err := json.NewDecoder(r.Body).Decode(&acrorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusBadRequest)
		return
	}

	if err := t.s.DeleteActorByID(r.Context(), acrorReq.Id); err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)
		return
	}
	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}
func (t transport) GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "GetActors"

	actors, err := t.s.GetActors(r.Context())
	if err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)
		return
	}
	t.handleOk(w, response.GetActors{Actors: actors}, methodName, http.StatusOK)

}
func (t transport) UpdateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "UpdateActor"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)

		return
	}
	var actorReq request.UpdateActor

	if err := json.NewDecoder(r.Body).Decode(&actorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusBadRequest)

		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(actorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusBadRequest)

		return
	}

	actor := models.Actor{
		FilmIds:   actorReq.Films,
		Name:      actorReq.Name,
		Gender:    actorReq.Gender,
		BirthDate: actorReq.BirthDate,

		Id: actorReq.Id,
	}
	if err := t.s.ChangeActor(r.Context(), actor); err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}
func (t transport) UpdateActorPartly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "UpdateActorPartly"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)

		return
	}
	var actorReq request.UpdateActor

	if err := json.NewDecoder(r.Body).Decode(&actorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusBadRequest)

		return
	}

	if err := validate.Struct(actorReq); err != nil {
		t.handleError(w, err, methodName, http.StatusBadRequest)

		return
	}

	actor := models.Actor{
		FilmIds:   actorReq.Films,
		Name:      actorReq.Name,
		Gender:    actorReq.Gender,
		BirthDate: actorReq.BirthDate,

		Id: actorReq.Id,
	}
	if err := t.s.ChangeActorPartly(r.Context(), actor); err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}
