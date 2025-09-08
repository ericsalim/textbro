package router

import (
	"github.com/ericsalim/textbro/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	pingHandler := handler.NewPingHandler()

	api := r.Group("/api")
	{
		api.GET("/ping", pingHandler.GetPing)
	}

	return r
}
