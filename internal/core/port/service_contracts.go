package ports

import "letsgo-flight-provider/internal/core/entities"

type FlightServiceContract interface {
	GetFlightList(source, destination, departure string) ([]entities.Flight, error)
	GetFlightById(id string) (entities.Flight, error)
	UpdateFlightById(id ,action string, count int) (bool, error)
	GetAircraftList() ([]string, error)
	GetcitytList() ([]string, error)
	GetListDaysWithFlight() ([]string, error)
}
