package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_application_routes(t *testing.T) {
	var register = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/web", "GET"},
	}

	var app application
	mux := app.routes()
	engine := mux.(*gin.Engine)

	for _, reg := range register {
		routeExists := CheckRouteExists(engine, reg.method, reg.route)
		assert.True(t, routeExists)
	}
}

// CheckRouteExists checks if a route exists in the router
func CheckRouteExists(router *gin.Engine, method string, path string) bool {
	for _, route := range router.Routes() {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}
