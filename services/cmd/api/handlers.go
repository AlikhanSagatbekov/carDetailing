package main

import (
	"net/http"

	"github.com/AlikhanSagatbekov/carDetailing/internal/data"
)

func (app *application) createServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service data.Service
	if err := app.readJSON(w, r, &service); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.db.Insert(&service); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"service": service}, nil)
}

func (app *application) getServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	service, err := app.db.Retrieve(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"service": service}, nil)
}

func (app *application) editServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var service data.Service
	if err := app.readJSON(w, r, &service); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	service.ID = id

	if err := app.db.Update(&service); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"service": service}, nil)
}

func (app *application) deleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.db.Delete(id); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Service deleted successfully"}, nil)
}
