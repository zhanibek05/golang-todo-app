package handler

import "github.com/gin-gonic/gin"

//

func (h *Handler) signUp(c *gin.Context) {

}
func (h *Handler) signIn(c *gin.Context) {
}
func (h *Handler) healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
