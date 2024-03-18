package main

import (
	"log"
	"vk/config"
	app "vk/internal/app"
)

// @title          Vk Intern
// @version         1.0
// @description     Тестовое задание

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Токен вида "Bearer access_token"
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
