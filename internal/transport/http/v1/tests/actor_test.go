package tests

import (
	"bytes"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
	"vk/internal/errs"
	"vk/internal/repository/models"
	mock_service "vk/internal/service/mocks"
	"vk/internal/transport/http/v1"

	"testing"
)

func TestFilm_AddActor(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"AddActor",
					[]any{mock.Anything, models.Actor{
						FilmIds:   []int{1, 2},
						Gender:    "female",
						BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Name:      "ivan",
					}},
					[]any{1, nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"id":1}`,
		},
		{
			name: "server err",

			inputBody: `{"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"AddActor",
					[]any{mock.Anything, models.Actor{
						FilmIds:   []int{1, 2},
						Gender:    "female",
						BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Name:      "ivan",
					}},
					[]any{-1, fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"birth_date: "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "not admin err",

			inputBody: `{"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{false, fmt.Errorf(errs.NotAdmin)},
				},
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NotAdmin),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {

				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("POST").Path("/actor/add").HandlerFunc(r.AddActor)
			req := httptest.NewRequest("POST", "/actor/add", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_UpdateActor(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id":1,"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeActor",
					[]any{mock.Anything, models.Actor{
						FilmIds:   []int{1, 2},
						Gender:    "female",
						BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Name:      "ivan",
						Id:        1,
					}},
					[]any{nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"status":"ok"}`,
		},
		{
			name: "server err",

			inputBody: `{"id":1,"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeActor",
					[]any{mock.Anything, models.Actor{
						FilmIds:   []int{1, 2},
						Gender:    "female",
						BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Name:      "ivan",
						Id:        1,
					}},
					[]any{fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"id","birth_date"2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "paginator err",

			inputBody: `{"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "not admin err",

			inputBody: `{"birth_date": "2000-01-01T00:00:00Z","films": [1,2],"gender": "female","name": "ivan"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{false, fmt.Errorf(errs.NotAdmin)},
				},
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NotAdmin),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {
				fmt.Println(mockFunc.methodName)
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("PUT").Path("/actor").HandlerFunc(r.UpdateActor)
			req := httptest.NewRequest("PUT", "/actor", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_UpdateActorPartly(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id":1,"films": [1,2],"gender": "female"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeActorPartly",
					[]any{mock.Anything, models.Actor{
						FilmIds: []int{1, 2},
						Gender:  "female",
						Id:      1,
					}},
					[]any{nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"status":"ok"}`,
		},
		{
			name: "server err",

			inputBody: `{"id":1,"films": [1,2],"gender": "female"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeActorPartly",
					[]any{mock.Anything, models.Actor{
						FilmIds: []int{1, 2},
						Gender:  "female",
						Id:      1,
					}},
					[]any{fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"id":1,"films: [1,2],"gender": "female"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "paginator err",

			inputBody: `{"films": [1,2],"gender": "female"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "not admin err",

			inputBody: `{"id":1,"films": [1,2],"gender": "female"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{false, fmt.Errorf(errs.NotAdmin)},
				},
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NotAdmin),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {
				fmt.Println(mockFunc.methodName)
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("PATCH").Path("/actor").HandlerFunc(r.UpdateActorPartly)
			req := httptest.NewRequest("PATCH", "/actor", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_DeleteActor(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id":1}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"DeleteActorByID",
					[]any{mock.Anything, 1},
					[]any{nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"status":"ok"}`,
		},
		{
			name: "server err",

			inputBody: `{"id":1}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"DeleteActorByID",
					[]any{mock.Anything, 1},
					[]any{fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"id":1"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "validation err id required",

			inputBody: `{"iddd":1"}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name: "not admin err",

			inputBody: `{"id":1}`,
			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{false, fmt.Errorf(errs.NotAdmin)},
				},
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NotAdmin),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("DELETE").Path("/actor").HandlerFunc(r.DeleteActor)
			req := httptest.NewRequest("DELETE", "/actor", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_GetActors(t *testing.T) {
	tests := []struct {
		name         string
		inputParams  string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:        "ok all",
			inputParams: ``,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetActors",
					[]any{mock.Anything},
					[]any{[]models.Actor{
						models.Actor{
							Id: 1,
							Films: []models.Film{
								{
									Title:       "test",
									Description: "12341234",
									Rating:      10,
									ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
								},
							},
							Gender:    "female",
							BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							Name:      "string",
						},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"actors":[{"id":1,"name":"string","gender":"female","birth_date":"2000-01-01T00:00:00Z","films":[{"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z"}]}]}`,
		},

		{
			name: "server err",

			inputParams: ``,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetActors",
					[]any{mock.Anything},
					[]any{nil, fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name:        "not logged err",
			inputParams: ``,
			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{false, nil},
				},
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NotLogged),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("GET").Path("/actor/all").HandlerFunc(r.GetActors)
			req := httptest.NewRequest("GET", "/actor/all"+test.inputParams, nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
