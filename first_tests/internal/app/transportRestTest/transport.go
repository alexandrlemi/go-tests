package transportRestTest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	ginItem *gin.Engine
	port    string
}

func NewTransport(port string) *Transport {
	r := gin.Default()

	return &Transport{ginItem: r, port: port}
}
func (t *Transport) Refresh(handler func(refToken string) error) error {

	// Обработчик для GET-запроса на путь "/hello"
	t.ginItem.GET("/refresh", func(c *gin.Context) {
		refToken := c.Query("refToken")
		handler(refToken)
	})
	return nil

}

func (t *Transport) Register(handler func(identifier, password string) error) error {
	t.ginItem.POST("/register", func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := handler(user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
		})
	})
	return nil
}

func (t *Transport) Run() error {

	// Запускаем сервер на порту 8080
	fmt.Println("Стартую на порту ", t.port)
	t.ginItem.Run(t.port)
	return nil
}
