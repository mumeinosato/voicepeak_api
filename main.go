package main

import (
	"log"

	"voicepeak_api/src"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	src.RegisterRoutes(router)

	log.Println("listening :3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
