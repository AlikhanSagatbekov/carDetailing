package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodPost, "/v1/services", app.createServiceHandler)
	router.HandlerFunc(http.MethodGet, "/v1/services/:id", app.getServiceHandler)
	router.HandlerFunc(http.MethodPut, "/v1/services/:id", app.editServiceHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/services/:id", app.deleteServiceHandler)

	return app.recoverPanic(app.rateLimit(router))
}
