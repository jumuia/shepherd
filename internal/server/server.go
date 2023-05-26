package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port       string
	DBName     string
	DBUser     string
	DBPassword string
	DBURL      string
	DBHost     string
	DBPort     string
}

func Run(cfg *Config) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    time.Now().String(),
		})
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	return r.Run(addr)
}
