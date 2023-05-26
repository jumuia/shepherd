package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jumuia/shepherd/internal/handler"
)

func setupRoutes(r *gin.Engine) {
	// Simple group: v1
	v1 := r.Group("/shepherd/api/v1")
	// User routes
	{
		userHandler := handler.NewUserHandler()
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users", userHandler.FetchAllUsers)
		v1.GET("/users/:user_id", userHandler.FetchUser)
		v1.PATCH("/users/:user_id", userHandler.UpdateUser)
		v1.DELETE("/users/:user_id", userHandler.DeleteUser)
	}
}
