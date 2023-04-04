package main

import (
	"encoding/json"
	"os"

	"/Users/mdvorsky/otra_carpeta/MODULO4/dia2/ej1t/cmd/server/handler"
	"/Users/mdvorsky/otra_carpeta/MODULO4/dia2/ej1t/internal/domain"
	"/Users/mdvorsky/otra_carpeta/MODULO4/dia2/ej1t/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/matidvo06/prueba.git/go/pkg/mod/github.com/gin-gonic/gin"
)

func main() {
	var productsList = []domain.Products{}
	loadProducts("products.json", &productsList)

	repo := product.NewRepository(productsList)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
	}
	r.Run(":8080")
}

// Carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
