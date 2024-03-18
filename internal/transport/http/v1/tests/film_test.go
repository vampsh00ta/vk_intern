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

func TestFilm_AddFilm(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"actors":[0],"description":"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"AddFilm",
					[]any{mock.Anything, models.Film{
						ActorIds:    []int{0},
						Description: "string",
						Rating:      10,
						ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Title:       "string",
					}},
					[]any{1, nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"id":1}`,
		},
		{
			name: "server err",

			inputBody: `{"actors":[0],"description":"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"AddFilm",
					[]any{mock.Anything, models.Film{
						ActorIds:    []int{0},
						Description: "string",
						Rating:      10,
						ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Title:       "string",
					}},
					[]any{-1, fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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

			inputBody: `{"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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
			router.Methods("POST").Path("/film/add").HandlerFunc(r.AddFilm)
			req := httptest.NewRequest("POST", "/film/add", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestFilm_UpdateFilm(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id":1,"actors":[0],"description":"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeFilm",
					[]any{mock.Anything, models.Film{
						Id:          1,
						ActorIds:    []int{0},
						Description: "string",
						Rating:      10,
						ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Title:       "string",
					}},
					[]any{nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"status":"ok"}`,
		},
		{
			name: "server err",

			inputBody: `{"id":1,"actors":[0],"description":"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeFilm",
					[]any{mock.Anything, models.Film{
						Id:          1,
						ActorIds:    []int{0},
						Description: "string",
						Rating:      10,
						ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						Title:       "string",
					}},
					[]any{fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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

			inputBody: `{"id":1,"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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
			router.Methods("PUT").Path("/film").HandlerFunc(r.UpdateFilm)
			req := httptest.NewRequest("PUT", "/film", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
func TestFilm_UpdateFilmPartly(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id":1,"actors":[0],"title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeFilmPartly",
					[]any{mock.Anything, models.Film{
						Id:       1,
						ActorIds: []int{0},
						Title:    "string",
					}},
					[]any{nil},
				},
			},
			expectedCode: 201,
			expectedBody: `{"status":"ok"}`,
		},
		{
			name: "server err",

			inputBody: `{"id":1,"actors":[0],"title":"string"}`,

			mockFuncs: []MockMethod{
				{
					"IsAdmin",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"ChangeFilmPartly",
					[]any{mock.Anything, models.Film{
						Id:       1,
						ActorIds: []int{0},
						Title:    "string",
					}},
					[]any{fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},

		{
			name: "validation err",

			inputBody: `{"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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
			name: "validation err rating greatere than 10",

			inputBody: `{"actors":[0],"description":"string","rating":11,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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

			inputBody: `{"id":1,"actors":[0],"description:"string","rating":10,"release_date":"2000-01-01T00:00:00Z","title":"string"}`,
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
			router.Methods("PATCH").Path("/film").HandlerFunc(r.UpdateFilmPartly)
			req := httptest.NewRequest("PATCH", "/film", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_DeleteFilm(t *testing.T) {
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
					"DeleteFilm",
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
					"DeleteFilm",
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
				fmt.Println(mockFunc.methodName)
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("DELETE").Path("/film").HandlerFunc(r.DeleteFilm)
			req := httptest.NewRequest("DELETE", "/film", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestFilm_GetFilms(t *testing.T) {
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
					"GetFilms",
					[]any{mock.Anything, "", ""},
					[]any{[]models.Film{
						models.Film{Id: 1,
							Title:       "12341234",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"films":[{"id":1,"title":"12341234","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z"}]}`,
		},
		{
			name:        "ok title search",
			inputParams: `?title=test`,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetFilmsByParams",
					[]any{mock.Anything, models.SortParams{
						ActorName: "",
						Title:     "test",
						SortBy:    "",
						OrderBy:   "",
					}},
					[]any{[]models.Film{
						models.Film{Id: 1,
							Title:       "test",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"films":[{"id":1,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z"}]}`,
		},
		{
			name:        "ok actor search",
			inputParams: `?name=test`,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetFilmsByParams",
					[]any{mock.Anything, models.SortParams{
						ActorName: "test",
						Title:     "",
						SortBy:    "",
						OrderBy:   "",
					}},
					[]any{[]models.Film{
						models.Film{
							Actors: []models.Actor{
								{
									Name: "test",
								},
							},

							Id:          1,
							Title:       "test",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"films":[{"id":1,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z","actors":[{"name":"test","birth_date":"0001-01-01T00:00:00Z"}]}]}`,
		},
		{
			name:        "ok actor+title search",
			inputParams: `?name=test&title=test`,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetFilmsByParams",
					[]any{mock.Anything, models.SortParams{
						ActorName: "test",
						Title:     "test",
						SortBy:    "",
						OrderBy:   "",
					}},
					[]any{[]models.Film{
						models.Film{
							Actors: []models.Actor{
								{
									Name: "test",
								},
							},

							Id:          1,
							Title:       "test",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"films":[{"id":1,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z","actors":[{"name":"test","birth_date":"0001-01-01T00:00:00Z"}]}]}`,
		},
		{
			name:        "ok  sort/order by",
			inputParams: `?sort_by=id&order_by=asc`,

			mockFuncs: []MockMethod{
				{
					"IsLogged",
					[]any{mock.Anything, mock.Anything},
					[]any{true, nil},
				},
				{
					"GetFilms",
					[]any{mock.Anything, "id", "asc"},
					[]any{[]models.Film{
						models.Film{Id: 1,
							Title:       "test",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
						models.Film{Id: 2,
							Title:       "test",
							Description: "12341234",
							Rating:      10,
							ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
					}, nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"films":[{"id":1,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z"},{"id":2,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z"}]}`,
		},
		//{
		//	name:        "ok  actor",
		//	inputParams: `?name=test`,
		//
		//	mockFuncs: []MockMethod{
		//		{
		//			"IsLogged",
		//			[]any{mock.Anything, mock.Anything},
		//			[]any{true, nil},
		//		},
		//		{
		//			"GetFilmsByActorName",
		//			[]any{mock.Anything, "test", "", ""},
		//			[]any{[]models.Film{
		//				models.Film{Id: 1,
		//
		//					Title:       "test",
		//					Description: "12341234",
		//					Rating:      10,
		//					ReleaseDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		//					Actors: []models.Actor{
		//						{Name: "test"},
		//					},
		//				},
		//			}, nil},
		//		},
		//	},
		//	expectedCode: 200,
		//	expectedBody: `{"films":[{"id":1,"title":"test","description":"12341234","rating":10,"release_date":"2000-01-01T00:00:00Z","actors":[{"name":"test","birth_date":"0001-01-01T00:00:00Z"}]}]}`,
		//},
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
					"GetFilms",
					[]any{mock.Anything, "", ""},
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
			router.Methods("GET").Path("/film/all").HandlerFunc(r.GetFilms)
			req := httptest.NewRequest("GET", "/film/all"+test.inputParams, nil)
			router.ServeHTTP(w, req)
			fmt.Println("/film/all" + test.inputParams)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
