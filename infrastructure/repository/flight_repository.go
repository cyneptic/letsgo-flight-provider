package repositories

import (
	"fmt"
	"letsgo-flight-provider/internal/core/entities"
	"time"
)

func (r *PGRepository) GetFlightList(source, destination, departureDateStr string) ([]entities.Flight, error) {
	var flights []entities.Flight

	dayStart, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", departureDateStr))
	if err != nil {
		return nil, err
	}
	dayEnd := dayStart.Add(24 * time.Hour)

	g := r.DB.Where("source = ? AND destination = ? AND departure_date >= ? AND departure_date < ?", source, destination, dayStart, dayEnd).Find(&flights)
	if g.Error != nil {
		return nil, g.Error
	}

	return flights, nil
}

func (r *PGRepository) GetFlightById(id string) (entities.Flight, error) {
	var flight entities.Flight

	g := r.DB.Where("id = ?", id).First(&flight)
	if g.Error != nil {
		return entities.Flight{}, g.Error
	}

	return flight, nil

}
