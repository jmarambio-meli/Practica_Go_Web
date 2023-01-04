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

	file, _ := ioutil.ReadFile("./db/db.json")

	_ = json.Unmarshal([]byte(file), &Productos)

	for i := 0; i < len(Productos); i++ {
		fmt.Println("Product", Productos[i].Name)
	}

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, Productos)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		for _, v := range Productos {
			if strconv.Itoa(v.Id) == c.Param("id") {
				c.JSON(200, v)
			}
		}
		c.String(404, "Producto Inexistente")
	})

	router.GET("/products/search", func(c *gin.Context) {
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
	})

	router.Run()
}
