package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/saludo", func(c *gin.Context) {
		var Persona persona

		//Modo para usar el package json -- obtengo el byte del body
		jsonData, err := c.GetRawData()
		if err != nil {
			fmt.Println(err)
		}

		//Con UnMarhsal
		if err := json.Unmarshal([]byte(string(jsonData)), &Persona); err != nil {
			log.Fatal(err)
		}

		//Con Decoder
		/*
			streaming := strings.NewReader(string(jsonData))
			decoder := json.NewDecoder(streaming)
			decoder.Decode(&Persona)
			for {
				if err := decoder.Decode(&Persona); err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("Hola %s %s", Persona.Nombre, Persona.Apellido)
				}
			}*/

		/* Manera mucho mas facil de hacer
		if err := c.BindJSON(&Persona); err != nil {
			return
		}
		fmt.Println(Persona)*/

		// Add the new album to the slice.
		c.String(http.StatusAccepted, fmt.Sprintf("Hola %s %s", Persona.Nombre, Persona.Apellido))
	})

	router.Run()
}
