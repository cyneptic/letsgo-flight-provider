package service

import (
	"letsgo-flight-provider/internal/core/entities"
	ports "letsgo-flight-provider/internal/core/port"
)

type FlightService struct {
	db ports.FlightRepositoryContract
}

func NewFlightService(db ports.FlightRepositoryContract) *FlightService {
	return &FlightService{
		db: db,
	}
}

func (svc *FlightService) GetFlightList(source, destination, departure string) ([]entities.Flight, error) {
	return svc.db.GetFlightList(source, destination, departure)
}

func (svc *FlightService) GetFlightById(id string) (entities.Flight, error) {
	return svc.db.GetFlightById(id)
}

func (svc *FlightService) UpdateFlightById(id, action string, count int) (bool, error) {
	return svc.db.UpdateFlightById(id, action, count)
}

func (svc *FlightService) GetAircraftList() ([]string, error) {
	return svc.db.GetAircraftList()
}
