package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var Productos []Producto

func main() {
	cargarJson("./db/db.json")

	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/products", getProducts)

	router.GET("/products/:id", getProductById)

	router.GET("/products/search", getProductByFilter)

	router.Run()
}

func cargarJson(url string) {
	file, _ := ioutil.ReadFile(url)
	_ = json.Unmarshal([]byte(file), &Productos)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func getProducts(c *gin.Context) {
	c.JSON(200, Productos)
}

func getProductById(c *gin.Context) {
	for _, v := range Productos {
		if strconv.Itoa(v.Id) == c.Param("id") {
			c.JSON(200, v)
		}
	}
	c.String(404, "Producto Inexistente")
}

func getProductByFilter(c *gin.Context) {
	valor := c.Query("valor")
	s, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		fmt.Println(err)
	}
	var productos []Producto
	for _, v := range Productos {

		if v.Price > s {
			productos = append(productos, v)
		}
	}
	c.JSON(200, productos)
}
