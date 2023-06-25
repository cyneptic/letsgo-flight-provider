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
	e.PATCH("/flights/:id", handler.UpdateFlightHandler)
	e.GET("/aircrafts", handler.ListAircraftsHandler)
	e.GET("/cities", handler.ListCitiesHandler)
	e.GET("/days-with-flight", handler.ListDaysWithFlightHandler)
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

func (h *FlightHandler) UpdateFlightHandler(c echo.Context) error {
	type RequestBody struct {
		Action string `json:"action"`
		Count  int    `json:"count"`
	}
	var requestBody RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest,
			"Invalid request body",
		)
	}

	action := requestBody.Action
	count := requestBody.Count

	err := validators.ValidateUpdateFlightParam(action)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	flight, err := h.svc.UpdateFlightById(c.Param("id"), action, count)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, flight)
}

func (h *FlightHandler) ListAircraftsHandler(c echo.Context) error {
	aircrafts, err := h.svc.GetAircraftList()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, aircrafts)
}

func (h *FlightHandler) ListCitiesHandler(c echo.Context) error {
	cities, err := h.svc.GetcitytList()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, cities)
}

func (h *FlightHandler) ListDaysWithFlightHandler(c echo.Context) error {
	days, err := h.svc.GetListDaysWithFlight()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, days)
}
