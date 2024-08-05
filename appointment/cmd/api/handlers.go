package main

import (
	"net/http"

	"github.com/AlikhanSagatbekov/carDetailing/internal/data"
)

func (app *application) createAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var appointment data.Appointment
	if err := app.readJSON(w, r, &appointment); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.db.Insert(&appointment); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"appointment": appointment}, nil)
}

func (app *application) getAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	appointment, err := app.db.Retrieve(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"appointment": appointment}, nil)
}

func (app *application) getExtendedAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	appointment, err := app.db.RetrieveExtended(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"appointment": appointment}, nil)
}

func (app *application) editAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var appointment data.Appointment
	if err := app.readJSON(w, r, &appointment); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	appointment.ID = id

	if err := app.db.Update(&appointment); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"appointment": appointment}, nil)
}

func (app *application) deleteAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if err := app.db.Delete(id); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Appointment deleted successfully"}, nil)
}
