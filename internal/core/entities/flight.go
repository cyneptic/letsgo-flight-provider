package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FlightNumber   string         `gorm:"size:255" json:"flight_number"`
	Source         string         `gorm:"size:255" json:"source"`
	Destination    string         `gorm:"size:255" json:"destination"`
	DepartureDate  time.Time      `json:"departure_date"`
	FlightDuration uint           `json:"flight_duration"`
	ArrivalDate    time.Time      `json:"arrival_date"`
	AirlineName    string         `gorm:"size:255" json:"airline_name"`
	AircraftName   string         `gorm:"size:255" json:"aircraft_name"`
	FareClass      FareClass      `gorm:"embedded"`
	Tax            uint64         `json:"tax"`
	FlightClass    string         `gorm:"size:255" json:"flight_class"`
	RemainingSeat  uint           `json:"remaining_seat"`
	CreatedAt      time.Time      `json:"created_at"`
	ModifiedAt     time.Time      `json:"modified_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type FareClass struct {
	AdultFare  uint64 `json:"adult_fare"`
	ChildFare  uint64 `json:"child_fare"`
	InfantFare uint64 `json:"infant_fare"`
}
