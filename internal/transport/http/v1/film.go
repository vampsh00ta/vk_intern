package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vk/internal/errs"
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

// @Summary     Add
// @Description Добавляет фильм. Возвратит ошибку,если указать id несуществующего актера.
// @Tags        Film
// @Accept      json
// @Param data body request.AddFilm true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.AddFilm
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /film/add [post]
func (t transport) AddFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "AddFilm"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, handleErrorAuth(err), methodName, http.StatusUnauthorized)

		return
	}

	var filmReq request.AddFilm

	if err := json.NewDecoder(r.Body).Decode(&filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	if err := validate.Struct(filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}
	film := models.Film{
		ActorIds:    filmReq.Actors,
		Title:       filmReq.Title,
		Rating:      filmReq.Rating,
		Description: filmReq.Description,
		ReleaseDate: filmReq.ReleaseDate,
	}
	id, err := t.s.AddFilm(r.Context(), film)
	if err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ServerError), methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.AddFilm{Id: id}, methodName, http.StatusCreated)

}

// @Summary     Delete
// @Description Удаляет фильм
// @Tags        Film
// @Accept      json
// @Param data body request.DeleteFilm true "Модель запроса"
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /film [delete]
func (t transport) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "DeleteFilm"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, handleErrorAuth(err), methodName, http.StatusUnauthorized)

		return
	}
	var filmReq request.DeleteFilm

	if err := json.NewDecoder(r.Body).Decode(&filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	if err := t.s.DeleteFilm(r.Context(), filmReq.Id); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ServerError), methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}

// @Summary     Get
// @Description Возвращает фильмы.
// @Tags        Film
// @Accept      json
// @Produce     json
// @Param name query string false "имя актера"
// @Param title query string false "название фильма"
// @Param sort_by query string false "поле сортировки"
// @Param order_by query string false "порядок сортировки"
// @Success     200 {object} response.GetFilms
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /film/all [get]
func (t transport) GetFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "GetFilms"
	if err := t.userPermission(w, r); err != nil {
		t.handleError(w, err, handleErrorAuth(err), methodName, http.StatusUnauthorized)
		return
	}
	var queryForm request.GetFilm
	//if err := schema.NewDecoder().Decode(&queryForm, r.Form); err != nil {
	//	t.handleError(w, err, methodName, http.StatusBadRequest)
	//
	//	return
	//}
	queryForm.SortBy = r.URL.Query().Get("sort_by")
	queryForm.OrderBy = r.URL.Query().Get("order_by")
	queryForm.Name = r.URL.Query().Get("name")
	queryForm.Title = r.URL.Query().Get("title")
	var films []models.Film
	var err error
	if queryForm.Title != "" || queryForm.Name != "" {
		films, err = t.s.GetFilmsByParams(r.Context(),
			models.SortParams{
				Title:     queryForm.Title,
				ActorName: queryForm.Name,
				SortBy:    queryForm.SortBy,
				OrderBy:   queryForm.OrderBy,
			},
		)

	} else {
		films, err = t.s.GetFilms(r.Context(), queryForm.SortBy, queryForm.OrderBy)

	}
	if err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ServerError), methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.GetFilms{Films: films}, methodName, http.StatusOK)

}

// @Summary     Update
// @Description Полностью изменяет фильмы
// @Tags        Film
// @Accept      json
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /film [put]
func (t transport) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "UpdateFilm"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, handleErrorAuth(err), methodName, http.StatusUnauthorized)

		return
	}
	var filmReq request.UpdateFilm

	if err := json.NewDecoder(r.Body).Decode(&filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	if err := validate.Struct(filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	film := models.Film{
		ActorIds:    filmReq.Actors,
		Title:       filmReq.Title,
		Rating:      filmReq.Rating,
		Description: filmReq.Description,
		ReleaseDate: filmReq.ReleaseDate,
		Id:          filmReq.Id,
	}
	if err := t.s.ChangeFilm(r.Context(), film); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ServerError), methodName, http.StatusInternalServerError)

		return
	}
	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}

// @Summary     UpdatePartly
// @Description Частично изменяет фильмы
// @Tags        Film
// @Accept      json
// @Produce     json
// @Success     201 {object} response.Ok
// @Failure     400 {object} response.Error
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Security ApiKeyAuth
// @Router      /film [patch]
func (t transport) UpdateFilmPartly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methodName := "UpdateFilmPartly"

	if err := t.adminPermission(w, r); err != nil {
		t.handleError(w, err, handleErrorAuth(err), methodName, http.StatusUnauthorized)

		return
	}
	var filmReq request.UpdateFilm

	if err := json.NewDecoder(r.Body).Decode(&filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	if err := validate.Struct(filmReq); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ValidationError), methodName, http.StatusBadRequest)

		return
	}

	film := models.Film{
		ActorIds:    filmReq.Actors,
		Title:       filmReq.Title,
		Rating:      filmReq.Rating,
		Description: filmReq.Description,
		ReleaseDate: filmReq.ReleaseDate,
		Id:          filmReq.Id,
	}
	if err := t.s.ChangeFilmPartly(r.Context(), film); err != nil {
		t.handleError(w, err, fmt.Errorf(errs.ServerError), methodName, http.StatusInternalServerError)

		return
	}

	t.handleOk(w, response.Ok{Status: "ok"}, methodName, http.StatusCreated)

}
