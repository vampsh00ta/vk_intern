package v1

import (
	"encoding/json"
	"net/http"
	_ "vk/cmd/vk_intern/docs"
	"vk/internal/repository/models"
	"vk/internal/transport/http/request"
	"vk/internal/transport/http/response"
)

// @Summary     Add
// @Description Создает актера
// @Tags        Actor
// @Accept      json
// @Param data body request.AddActor true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /actor/add [post]
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

// @Summary     Delete
// @Description Удаляет актера
// @Tags        Actor
// @Accept      json
// @Param data body request.DeleteActor true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /actor [delete]
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

// @Summary     Get
// @Description Возвращает актеров
// @Tags        Actor
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetActors
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /actor/all [get]
func (t transport) GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "GetActors"

	if err := t.userPermission(w, r); err != nil {
		t.handleError(w, err, methodName, http.StatusUnauthorized)
		return
	}
	actors, err := t.s.GetActors(r.Context())
	if err != nil {
		t.handleError(w, err, methodName, http.StatusInternalServerError)
		return
	}
	t.handleOk(w, response.GetActors{Actors: actors}, methodName, http.StatusOK)

}

// @Summary     Change
// @Description Полностью изменяет актера
// @Tags        Actor
// @Accept      json
// @Param data body request.UpdateActor true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /actor [put]
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

// @Summary     ChangePartly
// @Description Частично изменяет актера
// @Tags        Actor
// @Accept      json
// @Param data body request.UpdateActor true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /actor [patch]
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
