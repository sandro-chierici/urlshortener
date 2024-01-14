package routes

import (
	"log"
	"net/http"
	"strings"
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
func MapRoutes(r *gin.Engine, reg types.Register) {

	// Liveness probe for orchestrators
	r.GET("/healthz/liveness", getLiveness)

	// register an url and get shortened
	r.POST("/u", func(ctx *gin.Context) {
		shortendAndRegisterURL(ctx, reg)
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
func shortendAndRegisterURL(c *gin.Context, reg types.Register) {

	// getting url to register from the body
	var data urlBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, Response{Status: "ERROR", Err: "no url to register in request data"})
		return
	}

	log.Printf("Register URL %s", data.Url)

	code, err := reg.SetUrl(data.Url)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Status: "KO", Code: "", Err: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{Status: "OK", Code: code})
	}
}

// Get registered URL
func getShortened(c *gin.Context, reg types.Register) {

	code := c.Param("code")

	// get off "/"
	code = strings.Replace(code, "/", "", -1)
	if code == "" {
		c.JSON(http.StatusBadRequest, Response{Status: "KO", Code: "", Err: "Not found code into body"})
		return
	}

	url, err := reg.GetUrl(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Status: "KO", Code: "", Err: err.Error()})
	} else {
		c.JSON(http.StatusOK, Response{Status: "OK", Code: url})
	}
}
