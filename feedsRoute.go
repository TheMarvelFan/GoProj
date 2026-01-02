package main

import (
	"fmt"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedHandler(c *gin.Context, user database.User) {
	type fields struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := fields{}

	err := c.ShouldBindJSON(&params)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error binding json: %v", err))
		return
	}

	feed, createErr := apiCfg.DB.CreateFeeds(c, database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if createErr != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error creating feed: %v", createErr))
		return
	}

	sendJsonResponse(c, 201, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) getFeedsHandler(c *gin.Context) {
	feeds, err := apiCfg.DB.GetFeeds(c)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error fetching feeds: %v", err))
		return
	}

	sendJsonResponse(c, 200, dbFeedsToFeeds(feeds))
}
