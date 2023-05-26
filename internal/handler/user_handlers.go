package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	FetchUser(c *gin.Context)
	FetchAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

type userHandler struct {
}

// CreateUser implements UserHandler
func (*userHandler) CreateUser(c *gin.Context) {
	panic("unimplemented")
}

// DeleteUser implements UserHandler
func (*userHandler) DeleteUser(c *gin.Context) {
	panic("unimplemented")
}

type user struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// FetchAllUsers implements UserHandler
func (*userHandler) FetchAllUsers(c *gin.Context) {
	var users []user
	users = append(users, user{"alwin", 30})
	users = append(users, user{"jason", 30})
	users = append(users, user{"yesu", 30})
	users = append(users, user{"doss", 30})
	c.JSON(http.StatusOK, users)
}

// FetchUser implements UserHandler
func (*userHandler) FetchUser(c *gin.Context) {
	panic("unimplemented")
}

// UpdateUser implements UserHandler
func (*userHandler) UpdateUser(c *gin.Context) {
	panic("unimplemented")
}
