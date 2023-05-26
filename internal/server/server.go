package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jumuia/shepherd"
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
	subFS, err := fs.Sub(shepherd.WebappUIFS, "dist")
	if err != nil {
		return err
	}
	r.StaticFS("/webapp", http.FS(subFS))

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
