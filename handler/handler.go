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

	router.MaxMultipartMemory = 10 << 20

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.refresh)
	}

	userActions := router.Group("/user-actions")
	{
		userActions.POST("/fetch-search-results", h.search)
		userActions.POST("/fetch-news", h.getNews)
		userActions.POST("/like-post", h.likePost)
		userActions.POST("/follow", h.follow)
		userActions.POST("/get-liked-posts", h.getLikedPosts)
		userActions.POST("/get-user-posts", h.getUserPosts)
		userActions.POST("/get-user-info", h.getUserInfo)
		userActions.POST("/create-new-article", h.createNewArticle)
		userActions.POST("/change-settings", h.changeSettings)
		userActions.POST("/get-settings", h.getSettings)
	}

	return router
}
