package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/internal/domain"
	"github.com/jmarambio/prueba/internal/product"
	"github.com/jmarambio/prueba/pkg/web"
)

type ProductsHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductsHandler {
	return &ProductsHandler{service: service}
}

func (handler *ProductsHandler) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := handler.service.GetProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		web.SuccessReponse(c, 200, products)
	}
}

func (handler *ProductsHandler) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.FailureResponse(c, http.StatusBadRequest, errors.New("error en la conversion del ID"))
			return
		}
		product, err := handler.service.GetProductById(id)
		if err != nil {
			web.FailureResponse(c, http.StatusNotFound, errors.New("error al buscar el producto"))
			return
		}
		web.SuccessReponse(c, http.StatusOK, product)
	}
}

func (handler *ProductsHandler) GetProductByFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		valor := c.Query("valor")
		query, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, errors.New("error de transformacion de valor"))
			return
		}

		product, err := handler.service.GetProductByFilter(query)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, errors.New("error al buscar el producto"))
			return
		}
		web.SuccessReponse(c, http.StatusOK, product)
	}
}

func (handler *ProductsHandler) GetTotalByConsumer() gin.HandlerFunc {
	return func(c *gin.Context) {
		listaId := c.Query("list")
		ids := strings.Split(listaId, ",")
		var numbers []int
		for _, v := range ids {
			num, err := strconv.Atoi(v)
			if err != nil {
				web.FailureResponse(c, http.StatusInternalServerError, fmt.Errorf("error en el id: %s", v))
				return
			}
			numbers = append(numbers, num)
		}

		productos, total_price, err := handler.service.GetTotalByConsumer(numbers)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusOK, gin.H{
			"productos":   productos,
			"total_price": total_price,
		})
	}
}

func (handler *ProductsHandler) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.FailureResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		var producto domain.Producto

		if err := c.BindJSON(&producto); err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, errors.New("error en el body"))
			return
		}

		product, err := handler.service.AddProduct(producto)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusCreated, gin.H{
			"messages": "Producto a??adido exitosamente",
			"product":  product,
		})
	}
}

func (handler *ProductsHandler) EditProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.FailureResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.FailureResponse(c, http.StatusBadRequest, errors.New("error en la conversion del ID"))
			return
		}

		_, err = handler.service.GetProductById(id)
		if err != nil {
			web.FailureResponse(c, http.StatusNotFound, errors.New("error al buscar el producto"))
			return
		}

		var producto domain.Producto

		if err := c.BindJSON(&producto); err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, errors.New("error en el body"))
			return
		}

		product, err := handler.service.EditProduct(producto, id)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusOK, gin.H{
			"messages": "Producto editado exitosamente",
			"product":  product,
		})
	}
}

func (handler *ProductsHandler) PatchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.FailureResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.FailureResponse(c, http.StatusBadRequest, errors.New("error en la conversion del ID"))
			return
		}

		_, err = handler.service.GetProductById(id)
		if err != nil {
			web.FailureResponse(c, http.StatusNotFound, errors.New("error al buscar el producto"))
			return
		}

		var producto domain.Producto

		if err := c.BindJSON(&producto); err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, errors.New("error en el body"))
			return
		}

		product, err := handler.service.PatchProduct(producto, id)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusOK, gin.H{
			"messages": "Producto editado exitosamente con PATCH",
			"product":  product,
		})
	}
}

func (handler *ProductsHandler) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			web.FailureResponse(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.FailureResponse(c, http.StatusBadRequest, errors.New("error en la conversion del ID"))
			return
		}

		_, err = handler.service.GetProductById(id)
		if err != nil {
			web.FailureResponse(c, http.StatusNotFound, errors.New("error al buscar el producto"))
			return
		}

		err = handler.service.DeleteProduct(id)
		if err != nil {
			web.FailureResponse(c, http.StatusInternalServerError, err)
			return
		}
		web.SuccessReponse(c, http.StatusNoContent, "Producto Borrado Exitosamente")
	}
}
