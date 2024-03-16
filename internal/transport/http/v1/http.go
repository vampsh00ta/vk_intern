package v1

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"vk/internal/service"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type transport struct {
	s service.Service
	//l logger.Interface
}

func NewTransport(t service.Service) *http.ServeMux {
	r := &transport{t}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", r.Login)

	mux.HandleFunc("GET /film/all", r.GetFilms)
	mux.HandleFunc("POST /film", r.AddFilm)
	mux.HandleFunc("DELETE /film", r.DeleteFilm)
	mux.HandleFunc("PUT /film", r.UpdateFilm)
	mux.HandleFunc("PATCH /film", r.UpdateFilmPartly)

	mux.HandleFunc("GET /actor/all", r.GetActors)
	mux.HandleFunc("POST /actor", r.AddActor)
	mux.HandleFunc("DELETE /actor", r.DeleteActor)
	//mux.HandleFunc("PUT /actor", r.UpdateActor)
	//mux.HandleFunc("PATCH /actor", r.UpdateActor)

	return mux
	//h := handler.Group("/translation")
	//{
	//	h.GET("/history", r.history)
	//	h.POST("/do-translate", r.doTranslate)
	//}
}

//
////func NewRouter(handler *gin.Engine, l logger.Interface, t service.Service) {
////	// Options
////	handler.Use(gin.Logger())
////	handler.Use(gin.Recovery())
////
////	//// Swagger
////	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
////	//handler.GET("/swagger/*any", swaggerHandler)
////
////	//// K8s probe
////	//handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
////	//
////	//// Prometheus metrics
////	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
////
////	// Routers
////	h := handler.Group("/v1")
////	{
////		newTransport(h, t, l)
////	}
////}
