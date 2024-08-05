package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodPost, "/v1/appointments", app.createAppointmentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/appointments/:id", app.getAppointmentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/extendedAppointments/:id", app.getExtendedAppointmentHandler)
	router.HandlerFunc(http.MethodPut, "/v1/appointments/:id", app.editAppointmentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/appointments/:id", app.deleteAppointmentHandler)

	return app.recoverPanic(app.rateLimit(router))
}
