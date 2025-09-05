package main

import (
	"log"
	"net/http"

	"github.com/ericsalim/textbro/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}

	if cfg.ServerPort != "" {
		router.Run(":" + cfg.ServerPort)
	} else {
		router.Run()
	}

}
