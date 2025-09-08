package main

import (
	"log"

	"github.com/ericsalim/textbro/internal/config"
	"github.com/ericsalim/textbro/internal/router"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := router.SetupRouter()

	if cfg.ServerPort != "" {
		r.Run(":" + cfg.ServerPort)
	} else {
		r.Run()
	}

}
