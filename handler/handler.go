package handler

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
	router.Static("/avatars", "./avatars")
	router.Static("/images", "./images")
	
	userActions := router.Group("/user-actions")
	{
		userActions.POST("/fetch-search-results", h.search)
		userActions.POST("/fetch-news", h.getNews)
		userActions.POST("/like-post", h.likePost)
		userActions.POST("/follow", h.follow)
		userActions.POST("/get-liked-posts", h.getLikedPosts)
		userActions.POST("/get-user-posts", h.getUserPosts)
		userActions.POST("/get-user-info", h.getUserInfo)

		//userActions.POST("/change-settings")
	}

	return router
}
