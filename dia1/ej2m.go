/*
Ejercicio 2 - Manipulando el body
Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto 
“Hola + nombre + apellido”
El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:
{
		“nombre”: “Andrea”,
		“apellido”: “Rivas”
}
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context) {
		var persona Persona
		if err := c.BindJSON(&persona); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Petición inválida"})
			return
		}
		mensaje := "Hola " + persona.Nombre + " " + persona.Apellido
		c.JSON(http.StatusOK, gin.H{"mensaje": mensaje})
	})
	router.Run(":8080")
}

/*abrir POSTMAN y realizar una petición POST a la URK http://localhost:8080/saludo con la siguiente estructura de cuerpo:
{
	"nombre": "Andrea",
	"apellido": "Rivas"
}
Deberia recibir una respuesta en formato de texto que dice "Hola Andrea Rivas" */
