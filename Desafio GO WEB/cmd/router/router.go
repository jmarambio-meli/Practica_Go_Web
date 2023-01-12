package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine  *gin.Engine
	storage []domain.Ticket
}

func NewRouter(en *gin.Engine, storage []domain.Ticket) *Router {
	return &Router{engine: en, storage: storage}
}

func (r *Router) MapRoutes() {
	r.buildProductRoutes()
}

func (r *Router) buildProductRoutes() {
	repository := tickets.NewRepository(r.storage)
	service := tickets.NewService(repository)
	t := handler.NewTicketHandler(service)
	ticket := r.engine.Group("ticket")
	ticket.GET("/getByCountry/:dest", t.GetTicketsByCountry())
	ticket.GET("/getAverage/:dest", t.AverageDestination())

}
