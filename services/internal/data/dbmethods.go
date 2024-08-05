package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Insert(service *Service) error {
	_, err := m.DB.Exec("INSERT INTO services (name, description, price) VALUES ($1, $2, $3)",
		service.Name, service.Description, service.Price)
	if err != nil {
		return fmt.Errorf("failed to create a service: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (*Service, error) {
	var service Service
	row := m.DB.QueryRow("SELECT * FROM services WHERE id = $1", id)
	err := row.Scan(&service.ID, &service.Name, &service.Description, &service.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service: %w", err)
	}
	return &service, nil
}

func (m *DBModel) Update(service *Service) error {
	_, err := m.DB.Exec("UPDATE services SET  car_model = $1, color = $2, load = $3 WHERE id = $4",
		service.Name, service.Description, service.Price, service.ID)
	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM services WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}
	return nil
}
