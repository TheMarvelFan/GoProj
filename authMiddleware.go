package main

import (
	"fmt"

	"github.com/TheMarvelFan/GoPractice/internal/auth"
	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/gin-gonic/gin"
)

type authMiddleware func(*gin.Context, database.User)

func (apiCfg *apiConfig) authMiddlewareFunc(next authMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, err := auth.GetApiKey(c.Request.Header)

		if err != nil {
			sendErrorResponse(c, 403, fmt.Sprintf("Error extracting api key: %v", err))
			return
		}

		user, userErr := apiCfg.DB.GetUserByApiKey(c, apiKey)

		if userErr != nil {
			sendErrorResponse(c, 400, fmt.Sprintf("Error fetching user: %v", userErr))
			return
		}

		next(c, user)
	}
}
