package api

import (
	"cafe/config"
	"fmt"
	"github.com/gin-gonic/gin"
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
