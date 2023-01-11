package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/internal/product"
	"github.com/jmarambio/prueba/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServerProductsHandlerTest() *gin.Engine {
	_ = os.Setenv("TOKEN", "12345")

	storage := store.NewStore("./db.json")
	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	p := NewProductHandler(service)
	r := gin.Default()
	fmt.Println(storage)

	pr := r.Group("/products")
	pr.POST("/", p.AddProduct())
	pr.GET("/", p.GetProducts())
	pr.GET("/:id", p.GetProductById())
	pr.PUT("/:id", p.EditProduct())
	pr.PATCH("/:id", p.PatchProduct())
	pr.DELETE("/:id", p.DeleteProduct())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")
	return req, httptest.NewRecorder()
}

// PRACTICA 6 - EJERCICIO 1 TEST DE ÉXITO - TEST 1
func Test_GetProducts_OK(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodGet, "/products/", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

// PRACTICA 6 - EJERCICIO 1 TEST DE ÉXITO - TEST 2
func Test_GetProductById_OK(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodGet, "/products/200", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

// PRACTICA 6 - EJERCICIO 1 TEST DE ÉXITO - TEST 3
func Test_AddProduct_Created(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()
	request, response := createRequestTest(http.MethodPost, "/products/",
		` 
	{
		"name": "platano morado",
		"quantity": 2,
		"code_value": "S202027",
		"is_published": true,  
		"expiration": "26/03/2014",
		"price": 196.11
	}
	`,
	)
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code, "Error al crear el producto")
	assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"), "El Codigo debe ser 200")
}

// PRACTICA 6 - EJERCICIO 1 TEST DE ÉXITO - TEST 4
func Test_DeleteProduct_NoContent(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodDelete, "/products/500", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

// PRACTICA 6 - EJERCICIO 2 TEST DE FALLO - TEST 1
func Test_WrongId(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodGet, "/products/id_incorrecto", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	request, response = createRequestTest(http.MethodPut, "/products/id_incorrecto", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	request, response = createRequestTest(http.MethodPatch, "/products/id_incorrecto", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	request, response = createRequestTest(http.MethodDelete, "/products/id_incorrecto", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

// PRACTICA 6 - EJERCICIO 2 TEST DE FALLO - TEST 2
func Test_NotFoundId(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodGet, "/products/999", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)
	request, response = createRequestTest(http.MethodPut, "/products/999", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)
	request, response = createRequestTest(http.MethodPatch, "/products/999", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)
	request, response = createRequestTest(http.MethodDelete, "/products/999", "")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

// PRACTICA 6 - EJERCICIO 2 TEST DE FALLO - TEST 3
func Test_WrongToken(t *testing.T) {
	//Crea el servidor y define las rutas
	server := createServerProductsHandlerTest()

	request, response := createRequestTest(http.MethodPost, "/products/", ` 
	{
		"name": "platano morado",
		"quantity": 2,
		"code_value": "S202027",
		"is_published": true,  
		"expiration": "26/03/2014",
		"price": 196.11
	}
	`)
	request.Header.Set("token", "tokeninvalido")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	request, response = createRequestTest(http.MethodPut, "/products/999", "")
	request.Header.Set("token", "tokeninvalido")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	request, response = createRequestTest(http.MethodPatch, "/products/999", "")
	request.Header.Set("token", "tokeninvalido")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	request, response = createRequestTest(http.MethodDelete, "/products/999", "")
	request.Header.Set("token", "tokeninvalido")
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnauthorized, response.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &req.Response.Request)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}
