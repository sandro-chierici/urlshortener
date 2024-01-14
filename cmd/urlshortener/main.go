package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortener/v2/internal/encoders"
	"urlshortener/v2/internal/register"
	"urlshortener/v2/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// set listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// shortener algo implementation
	var encoder = encoders.New()

	// concrete register implementation (in memory, redis isn't ready)
	var reg = register.NewSimpleRegister(encoder)

	routes.MapRoutes(router, reg)

	http.ListenAndServe(fmt.Sprint(":", port), router)
}
