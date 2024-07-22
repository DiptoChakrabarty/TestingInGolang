package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	mux := gin.Default()
	mux.LoadHTMLGlob("./templates/*")
	// register middleware
	mux.Use(gin.Recovery())

	// go to routes
	mux.GET("/", app.Home)

	return mux
}
