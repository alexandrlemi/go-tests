package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main()  {
	r := gin.Default()

	// Определяем маршрут для GET-запроса на корневой путь "/"
	r.GET("/",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Привет, мир!",
			})
	})

	// Определяем маршрут для GET-запроса на путь "/hello/:name"
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name") // Получаем параметр "name" из URL
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, " + name + "!",
		})
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
	
}