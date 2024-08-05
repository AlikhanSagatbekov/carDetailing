package data

type Appointment struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	ServiceID  int    `json:"service_id"`
	Date       string `json:"date"`
	Status     string `json:"status"`
}

type ExtendedAppointment struct {
	ID                 int     `json:"id"`
	CustomerID         int     `json:"customer_id"`
	ServiceID          int     `json:"service_id"`
	Date               string  `json:"date"`
	Status             string  `json:"status"`
	ServiceName        string  `json:"service_name"`
	ServiceDescription string  `json:"service_description"`
	ServicePrice       float64 `json:"service_price"`
}
