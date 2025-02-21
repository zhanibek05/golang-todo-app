package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/id", h.getListById)
			lists.PUT("/id", h.updateList)
			lists.DELETE("/id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("id/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/item_id", h.getItemById)
				items.PUT("item_id/", h.updateItem)
				items.DELETE("item_id/", h.deleteItem)
			}
		}
	}

	return router
}
