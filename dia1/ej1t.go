/*
Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un 
servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:
https://gyazo.com/d3913068815e93b1bd38a82f5092da7e
El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos de productos. Esta slice se debe 
cargar cada vez que se inicie la API para realizar las distintas consultas.
Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor utilizando el paquete gin en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.  
Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
Crear una ruta /products/:id que nos devuelva un producto por su id.
Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	file, err := os.Open("productos.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var products []Product

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, p := range products {
			if strconv.Itoa(p.ID) == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.String(http.StatusNotFound, "Producto no Encontrado")
	})

	router.GET("/products/search", func(c *gin.Context) {
		priceGtStr := c.Query("priceGt")
		priceGt, err := strconv.ParseFloat(priceGtStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Parámetro inválido")
			return
		}

		var result []Product
		for _, p := range products {
			if p.Price > priceGt {
				result = append(result, p)
			}
		}

		c.JSON(http.StatusOK, result)
	})

	//Iniciar el servidor en el puerto 8080
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
