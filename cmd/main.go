package main

import (
	"post_api_gateway/api"
	"post_api_gateway/config"
	"post_api_gateway/pkg/logger"
	"post_api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "post_api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
