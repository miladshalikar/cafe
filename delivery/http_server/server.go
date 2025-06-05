package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miladshalikar/cafe/config"
	"net/http"
)

func InitServer(cfg config.ServerConfig) {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
