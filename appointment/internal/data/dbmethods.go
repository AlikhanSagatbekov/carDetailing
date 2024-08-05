package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Insert(appointment *Appointment) error {
	_, err := m.DB.Exec("INSERT INTO appointments (customer_id, service_id, date, status) VALUES ($1, $2, $3, $4)",
		appointment.CustomerID, appointment.ServiceID, appointment.Date, appointment.Status)
	if err != nil {
		return fmt.Errorf("failed to create an appointment: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (*Appointment, error) {
	var appointment Appointment
	row := m.DB.QueryRow("SELECT * FROM appointments WHERE id = $1", id)
	err := row.Scan(&appointment.ID, &appointment.CustomerID, &appointment.ServiceID, &appointment.Date, &appointment.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve an appointment: %w", err)
	}
	return &appointment, nil
}

func (m *DBModel) RetrieveExtended(id int) (*ExtendedAppointment, error) {
	var extendedAppointment ExtendedAppointment
	row := m.DB.QueryRow("SELECT a.id AS appointment_id, a.customer_id, s.id AS service_id, a.date, a.status, s.name AS service_name, s.description AS service_description, s.price AS service_price FROM appointments a JOIN services s ON a.service_id = s.id WHERE a.id =$1", id)
	err := row.Scan(&extendedAppointment.ID, &extendedAppointment.CustomerID, &extendedAppointment.ServiceID, &extendedAppointment.Date, &extendedAppointment.Status, &extendedAppointment.ServiceName, &extendedAppointment.ServiceDescription, &extendedAppointment.ServicePrice)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve an appointment: %w", err)
	}
	return &extendedAppointment, nil
}

func (m *DBModel) Update(appointment *Appointment) error {
	_, err := m.DB.Exec("UPDATE appointments SET customer_id = $1, service_id = $2, date = $3, status = $4 WHERE id = $5",
		appointment.CustomerID, appointment.ServiceID, appointment.Date, appointment.Status, appointment.ID)
	if err != nil {
		return fmt.Errorf("failed to update an appointment: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM appointments WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete an appointment: %w", err)
	}
	return nil
}
