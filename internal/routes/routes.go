package routes

import (
	"log"
	"net/http"
	"urlshortener/v2/internal/shortener"
	"urlshortener/v2/internal/types"

	"github.com/gin-gonic/gin"
)

// Generic OK response
type Response struct {
	Status string
	Code   string
	Err    string
}

type urlBody struct {
	Url string
}

// Mapping of exposed endpoints
func MapRoutes(r *gin.Engine, reg types.Register, sh *shortener.SimpleShortener) {

	r.GET("/healthz/liveness", getLiveness)

	r.POST("/u", func(ctx *gin.Context) {
		shortendAndRegisterURL(ctx, reg, sh)
	})

	r.GET("/u/:code", func(ctx *gin.Context) {
		getShortened(ctx, reg)
	})
}

// Liveness probe
func getLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Status: "OK"})
}

// register	new url
func shortendAndRegisterURL(c *gin.Context, reg types.Register, sh *shortener.SimpleShortener) {

	// getting url to register from the body
	var data urlBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, Response{Status: "ERROR", Err: "no url to register in request data"})
		return
	}

	log.Printf("Register URL %s", data.Url)

	// calculate shortened code
	code := sh.Evaluate(data.Url)

	// for speed concern here I reply always ok
	// and in parallel let's check cache for existance
	go func() {

		_, err := reg.GetShortened(code)
		if err == nil {
			// exists
			return
		}

		// register in cacche
		err = reg.SetUrl(data.Url, code)
		if err != nil {
			log.Printf("Error registering URL %s: %s", data.Url, err)
			return
		}
	}()

	c.JSON(http.StatusOK, Response{Status: "OK", Code: code})
}

// Get registered URL
func getShortened(c *gin.Context, reg types.Register) {

	code := c.Param("code")
	full, err := reg.GetShortened(code)
	if err == nil {
		// exists
		c.JSON(http.StatusOK, Response{Status: "OK", Code: full})
		return
	}
	c.JSON(http.StatusBadRequest, Response{Status: "ERROR", Err: "Not Found"})
}
