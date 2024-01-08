package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortener/v2/internal/register"
	"urlshortener/v2/internal/routes"
	"urlshortener/v2/internal/shortener"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// set listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// concrete register implementation (Redis)
	var reg = register.NewRedisRegister()
	// shortener algo implementation
	var shortener = shortener.New()

	routes.MapRoutes(router, reg, shortener)

	http.ListenAndServe(fmt.Sprint(":", port), router)
}
