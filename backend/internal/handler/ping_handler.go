package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewPingHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
