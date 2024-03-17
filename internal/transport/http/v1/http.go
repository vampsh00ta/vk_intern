package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
	_ "vk/cmd/vk_intern/docs"
	"vk/internal/service"
	//swaggerFiles "github.com/swaggo/files"
	// This is what I needed to add for it to work, this is "docs" in the root of my application
	// generated with swag init
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type transport struct {
	s service.Service
	l *zap.SugaredLogger
}

func NewTransport(t service.Service, l *zap.SugaredLogger) http.Handler {
	r := &transport{t, l}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))
	mux.HandleFunc("POST /login", r.Login)

	mux.HandleFunc("GET /film/all", r.GetFilms)
	mux.HandleFunc("POST /film/add", r.AddFilm)
	mux.HandleFunc("DELETE /film", r.DeleteFilm)
	mux.HandleFunc("PUT /film", r.UpdateFilm)
	mux.HandleFunc("PATCH /film", r.UpdateFilmPartly)

	mux.HandleFunc("GET /actor/all", r.GetActors)
	mux.HandleFunc("POST /actor/add", r.AddActor)
	mux.HandleFunc("DELETE /actor", r.DeleteActor)
	mux.HandleFunc("PUT /actor", r.UpdateActor)
	mux.HandleFunc("PATCH /actor", r.UpdateActorPartly)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:8080/"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	return handler

}
