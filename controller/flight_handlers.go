package controllers

import (
	"letsgo-flight-provider/controller/validators"
	repositories "letsgo-flight-provider/infrastructure/repository"
	"letsgo-flight-provider/internal/core/entities"
	ports "letsgo-flight-provider/internal/core/port"
	"letsgo-flight-provider/internal/core/service"
	"net/http"

	"github.com/labstack/echo"
)

type FlightHandler struct {
	svc ports.FlightServiceContract
}

func NewFlightHandler(svc ports.FlightServiceContract) *FlightHandler {
	return &FlightHandler{
		svc: svc,
	}
}

func AddFlightRoutes(e *echo.Echo) {
	db := repositories.NewGormDatabase()
	svc := service.NewFlightService(db)
	handler := NewFlightHandler(svc)
	e.GET("/flights", handler.ListFlightsHandler)
	e.GET("/flights/:id", handler.FindFlightHandler)

}
func (h *FlightHandler) ListFlightsHandler(c echo.Context) error {
	var flightList []entities.Flight
	err := validators.ValidateListFlightParam(c.QueryParams())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	flightList, err = h.svc.GetFlightList(c.QueryParam("source"), c.QueryParam("destination"), c.QueryParam("departing"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, flightList)
	
}

func (h *FlightHandler) FindFlightHandler(c echo.Context) error {
	var flight entities.Flight
	flight, err := h.svc.GetFlightById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, flight)
}