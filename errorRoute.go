package main

import "github.com/gin-gonic/gin"

func errorRouteHandler(c *gin.Context) {
	sendErrorResponse(c, 400, "Sum Ting Wong")
}
