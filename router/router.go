package router

import (
	"web_demo/handler/home"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	homeGroup := g.Group("/home")
	{
		homeGroup.GET("/name", home.HomeName)
	}

	return g
}
