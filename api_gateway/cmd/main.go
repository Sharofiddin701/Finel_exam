package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"api_gateway/api"
	"api_gateway/api/handlers"
	"api_gateway/config"
	"api_gateway/pkg/logger"
	"api_gateway/services"
)

func main() {
	cfg := config.Load()

	grpcSvcs, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("ui_go_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
