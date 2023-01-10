package handler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jmarambio/prueba/internal/domain"
)

func CargarJson(url string) {
	file, _ := ioutil.ReadFile(url)
	_ = json.Unmarshal([]byte(file), &domain.Productos)
}

/*
	func GetProducts(c *gin.Context) {
		service := product.NewService()
		products, err := service.List()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al recuperar la lista de productos")
			return
		}
		c.JSON(http.StatusOK, products)
	}

	func GetProductById(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error en la conversion del ID")
			return
		}
		service := product.NewService()
		product, err := service.ListId(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al buscar el producto")
			return
		}
		c.JSON(http.StatusOK, product)
	}

	func GetProductByFilter(c *gin.Context) {
		valor := c.Query("valor")
		query, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println(err)
		}

		service := product.NewService()
		product, err := service.ListFilter(query)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al buscar el producto")
			return
		}
		c.JSON(http.StatusOK, product)

}

	func AddProduct(c *gin.Context) {
		var producto domain.Producto

		if err := c.BindJSON(&producto); err != nil {
			c.String(http.StatusInternalServerError, "Error en el body")
			return
		}

		service := product.NewService()
		product, err := service.AddList(producto)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": "Producto agregado exitosamente",
			"product":  product,
		})
	}
*/
