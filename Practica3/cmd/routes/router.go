package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/cmd/handler"
	"github.com/jmarambio/prueba/internal/domain"
	"github.com/jmarambio/prueba/internal/product"
)

type Router struct {
	engine *gin.Engine
	//group     *gin.RouterGroup
	productos []domain.Producto
}

func NewRouter(en *gin.Engine, productos []domain.Producto) *Router {
	return &Router{engine: en, productos: productos}
}

/*
func (r *Router) MapRoutes() {
	r.group.Group("/api/v1")
	r.buildPingRoute()
}

func (r *Router) buildPingRoute() {
	r.group.GET("ping", nil)
}*/

func (r *Router) buildProductRoutes() {
	repository := product.NewRepository(r.productos)
	service := product.NewService(repository)
	p := handler.NewProductHandler(*service)
	product := r.engine.Group("products")
	product.GET("", p.GetProducts())
	product.GET("/:id", p.GetProductById())
	product.GET("/search", p.GetProductByFilter())
	product.POST("", p.AddProduct())
}

func (r *Router) SetRoutes() {
	r.buildProductRoutes()
}
