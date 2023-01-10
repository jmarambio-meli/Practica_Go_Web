package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
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
	products := router.Group("/products")
	{
		products.GET("", getProducts)
		products.GET(":id", getProductById)
		products.GET("/search", getProductByFilter)
		products.POST("", addProduct)
	}

	router.Run()
}

func cargarJson(url string) {
	file, _ := ioutil.ReadFile(url)
	_ = json.Unmarshal([]byte(file), &Productos)
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

func addProduct(c *gin.Context) {
	var producto Producto
	if err := c.BindJSON(&producto); err != nil {
		return
	}

	err := validaciones(producto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = codeValueRepeated(producto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	producto.Id = len(Productos) + 1
	Productos = append(Productos, producto)
	c.JSON(200, gin.H{
		"messages": "Producto agregado exitosamente",
		"product":  producto,
	})
}

func codeValueRepeated(p Producto) (bool, error) {
	for _, v := range Productos {
		if v.Code_value == p.Code_value {
			return true, errors.New("el code value ya existe")
		}
	}
	return false, nil
}

func validaciones(p Producto) error {
	if p.Name == "" {
		return errors.New("el nombre no puede estar vacio")
	}
	if p.Expiration == "" {
		return errors.New("la fecha de expiracion no puede estar vacio")
	}
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if !re.MatchString(p.Expiration) {
		return errors.New("fecha incorrecta o Fomato incorrecto de expiración, el formato es : dd/mm/yyyy")
	}
	if p.Code_value == "" {
		return errors.New("el code value no puede ser vacio")
	}
	if p.Price <= 0 {
		return errors.New("el precio no puede ser igual o menor a 0")
	}
	if p.Quantity <= 0 {
		return errors.New("la cantidad no puede ser igual o menor a 0")
	}
	return nil
}
