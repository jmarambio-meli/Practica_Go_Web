package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/bootcamp-go/desafio-go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type Ticket struct {
	service tickets.Service
}

func NewTicketHandler(s tickets.Service) *Ticket {
	return &Ticket{
		service: s,
	}
}

func (t *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := t.service.GetTotalTickets(c, destination)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusOK, tickets)
	}
}

func (t *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := t.service.AverageDestination(c, destination)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}

		web.SuccessReponse(c, http.StatusOK, avg)
	}
}
