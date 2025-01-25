package main

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	userActions := router.Group("/user-actions")
	{
		userActions.GET("/fetch-people", h.fetchPeople)
	}

	return router
}
