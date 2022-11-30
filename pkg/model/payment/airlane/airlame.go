package airlane

import "time"

// AirlineTicket Информация о пассажирах и билетах передается при создании платежа.
// В запросе обязательно указывается или номер билета (ticket_number),
// или номер брони (booking_reference), если номера билета пока нет.
type AirlineTicket struct {
	// Уникальный номер билета. Если при создании платежа вы уже знаете номер билета,
	// тогда ticket_number — обязательный параметр. Если не знаете,
	// тогда вместо ticket_number необходимо передать booking_reference с номером бронирования.
	TicketNumber string `json:"ticket_number,omitempty"`
	// Номер бронирования. Обязателен, если не передан ticket_number
	BookingReference string     `json:"booking_reference,omitempty"`
	Passengers       Passengers `json:"passengers,omitempty"`
	Legs             Legs       `json:"legs,omitempty"`
}

type Passenger struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// Passengers Список пассажиров.
type Passengers []*Passenger

type Leg struct {
	DepartureAirport   string    `json:"departure_airport,omitempty"`
	DestinationAirport string    `json:"destination_airport,omitempty"`
	DepartureDate      time.Time `json:"departure_date,omitempty"`
	CarrierCode        string    `json:"carrier_code,omitempty"`
}

// Legs Список перелетов.
type Legs []*Leg
