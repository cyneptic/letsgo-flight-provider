package ports

import "letsgo-flight-provider/internal/core/entities"

type FlightRepositoryContract interface {
	GetFlightList(source, destination, departure string) ([]entities.Flight, error)
}
