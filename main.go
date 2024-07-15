package main

import (
	"github.com/AwahMarwah/Technical_test/controller"
	"github.com/AwahMarwah/Technical_test/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fruitService := services.NewFruitService()
	fruitHandler := controller.NewFruitHandler(fruitService)

	r.GET("/fruits", fruitHandler.GetAllFruitNames)
	r.GET("/fruits/separated", fruitHandler.SeparateFruitsByType)
	r.GET("/fruits/stock", fruitHandler.GetTotalStockByType)
	r.GET("/comments", fruitHandler.Comments)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
