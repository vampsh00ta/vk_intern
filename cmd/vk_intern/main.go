package main

import (
	"log"
	"vk/config"
	app "vk/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	//r, err := rand.Int(rand.Reader, big.NewInt(2))
	//fmt.Println(r, err)
	app.Run(cfg)

}
