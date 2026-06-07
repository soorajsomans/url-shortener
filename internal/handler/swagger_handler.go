package handler

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwaggerRoutes(
	mux *http.ServeMux,
) {
	mux.Handle(
		"/swagger/",
		httpSwagger.WrapHandler,
	)
}
