package tests

import (
	"bytes"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"strings"
	"testing"
	"vk/internal/errs"
	mock_service "vk/internal/service/mocks"
	v1 "vk/internal/transport/http/v1"
)

func TestFilm_Login(t *testing.T) {
	tests := []struct {
		name         string
		inputBody    string
		mockFuncs    []MockMethod
		expectedCode int
		expectedBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username":"admin"}`,

			mockFuncs: []MockMethod{
				{
					"Login",
					[]any{mock.Anything, "admin"},
					[]any{"jwt_token", nil},
				},
			},
			expectedCode: 200,
			expectedBody: `{"access":"jwt_token"}`,
		},
		{
			name:      "no such userc",
			inputBody: `{"username":"test"}`,

			mockFuncs: []MockMethod{
				{
					"Login",
					[]any{mock.Anything, "test"},
					[]any{"", fmt.Errorf(errs.NoUserSuchUser)},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.NoUserSuchUser),
		},
		{
			name:      "some error",
			inputBody: `{"username":"test"}`,

			mockFuncs: []MockMethod{
				{
					"Login",
					[]any{mock.Anything, "test"},
					[]any{"", fmt.Errorf("some error")},
				},
			},
			expectedCode: 500,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ServerError),
		},
		{
			name:      "null username error",
			inputBody: `{"usefffrname":"test"}`,

			mockFuncs:    []MockMethod{},
			expectedCode: 400,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
		{
			name:      "validation error",
			inputBody: `{"usefffrname:"test"}`,

			mockFuncs:    []MockMethod{},
			expectedCode: 400,
			expectedBody: fmt.Sprintf(`{"error":"%s"}`, errs.ValidationError),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			srvc := mock_service.NewService(t)
			for _, mockFunc := range test.mockFuncs {
				fmt.Errorf(mockFunc.methodName)
				srvc.On(mockFunc.methodName, mockFunc.args...).Once().Return(mockFunc.returns...)
			}

			r := v1.TestTrasport(srvc, LoadLoggerDev())
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Methods("POST").Path("/login").HandlerFunc(r.Login)
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)
			assert.Equal(t, test.expectedCode, w.Code)
			assert.Equal(t, test.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
