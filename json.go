package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(c *gin.Context, code int, errMsg string) {
	c.Error(errors.New(errMsg))
	sendJsonResponse(c, code, gin.H{
		"error": errMsg,
	})
}

func sendJsonResponse(c *gin.Context, code int, data any) {
	if data == nil {
		c.JSON(code, gin.H{})
		return
	}

	c.JSON(code, data)
}
