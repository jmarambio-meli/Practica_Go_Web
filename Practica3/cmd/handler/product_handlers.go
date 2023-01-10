package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/internal/domain"
	"github.com/jmarambio/prueba/internal/product"
)

type ProductsHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductsHandler {
	return &ProductsHandler{service: service}
}

func (handler *ProductsHandler) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := handler.service.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(200, products)
	}
}

func (handler *ProductsHandler) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error en la conversion del ID")
			return
		}
		product, err := handler.service.ListId(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al buscar el producto")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func (handler *ProductsHandler) GetProductByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		valor := c.Query("valor")
		query, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println(err)
		}

		product, err := handler.service.ListFilter(query)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al buscar el producto")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func (handler *ProductsHandler) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var producto domain.Producto

		if err := c.BindJSON(&producto); err != nil {
			c.String(http.StatusInternalServerError, "Error en el body")
			return
		}

		product, err := handler.service.AddList(producto)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": "Producto agregado exitosamente",
			"product":  product,
		})
	}
}
