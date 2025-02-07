package transportRestTest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	ginItem *gin.Engine
	port string
}

func NewTransport(port string) *Transport{
	r := gin.Default()

	return &Transport{ginItem: r,port: port}
}
func (t *Transport) Refrash(handler func(refToken string) error) error {

	// Обработчик для GET-запроса на путь "/hello"
	t.ginItem.GET("/refrash", func(c *gin.Context) {
		refToken := c.Query("refToken") 
		handler(refToken)
	})
	return nil

}
func (t *Transport) Run()  error {
	
	// Запускаем сервер на порту 8080
	fmt.Println("Стартую на порту ",t.port)
	t.ginItem.Run(t.port)
	return nil
}