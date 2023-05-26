package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:5050"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    time.Now().String(),
		})
	})
	setupRoutes(r)

	addr := fmt.Sprintf(":%s", cfg.Port)
	return r.Run(addr)
}
