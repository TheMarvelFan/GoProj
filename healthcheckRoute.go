package main

import "github.com/gin-gonic/gin"

func healthCheckRouteHandler(c *gin.Context) {
	sendJsonResponse(c, 200, gin.H{
		"status": "OK",
	})
}
