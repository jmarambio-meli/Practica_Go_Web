package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/cmd/handler"
	"github.com/jmarambio/prueba/internal/product"
	"github.com/jmarambio/prueba/pkg/store"
)

type Router struct {
	engine *gin.Engine
	//group     *gin.RouterGroup
	storage store.Store
}

func NewRouter(en *gin.Engine, storage store.Store) *Router {
	return &Router{engine: en, storage: storage}
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
	repository := product.NewRepository(r.storage)
	service := product.NewService(repository)
	p := handler.NewProductHandler(service)
	product := r.engine.Group("products")
	product.GET("", p.GetProducts())
	product.GET("/:id", p.GetProductById())
	product.GET("/search", p.GetProductByFilter())
	product.GET("/consumer_price", p.GetTotalByConsumer())
	product.POST("", p.AddProduct())
	product.PUT("/:id", p.EditProduct())
	product.PATCH("/:id", p.PatchProduct())
	product.DELETE("/:id", p.DeleteProduct())
}

func (r *Router) SetRoutes() {
	r.buildProductRoutes()
}
