package main

import (
	"fmt"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createFeedFollowersHandler(c *gin.Context, user database.User) {
	type fields struct {
		FeedId uuid.UUID `json:"feedId"`
	}

	params := fields{}

	err := c.ShouldBindJSON(&params)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error binding json: %v", err))
		return
	}

	feedFollower, createErr := apiCfg.DB.CreateFeedFollowers(c, database.CreateFeedFollowersParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if createErr != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error creating feed follower: %v", createErr))
		return
	}

	sendJsonResponse(c, 201, dbFeedFollowerToFeedFollower(feedFollower))
}

func (apiCfg *apiConfig) getFeedsFollowedByUserHandler(c *gin.Context, user database.User) {
	feedFollowings, err := apiCfg.DB.GetFeedsFollowedByUser(c, user.ID)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error getting feeds followed by user: %v", err))
		return
	}

	feedsFollowed := []database.Feed{}

	for _, feedFollowing := range feedFollowings {
		feed, feedErr := apiCfg.DB.GetFeedByID(c, feedFollowing.FeedID)

		if feedErr != nil {
			sendErrorResponse(c, 400, fmt.Sprintf("Error fetching feed: %v", feedErr))
		}

		feedsFollowed = append(feedsFollowed, feed)
	}

	sendJsonResponse(c, 200, dbFeedsToFeeds(feedsFollowed))
}

func (apiCfg *apiConfig) getFollowersForFeedHandler(c *gin.Context, user database.User) {
	type fields struct {
		FeedId uuid.UUID `json:"feedId"`
	}

	params := fields{}

	err := c.ShouldBindJSON(&params)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error binding json: %v", err))
		return
	}

	feedFollowings, err := apiCfg.DB.GetFollowersForFeed(c, params.FeedId)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error getting followers for feed: %v", err))
		return
	}

	feedFollowers := []database.User{}

	for _, feedFollowing := range feedFollowings {
		user, userErr := apiCfg.DB.GetUserByID(c, feedFollowing.UserID)

		if userErr != nil {
			sendErrorResponse(c, 400, fmt.Sprintf("Error fetching user: %v", userErr))
		}

		feedFollowers = append(feedFollowers, user)
	}

	sendJsonResponse(c, 200, dbUsersToUsers(feedFollowers))
}

func (apiCfg *apiConfig) deleteFeedFollowingHandler(c *gin.Context, user database.User) {
	feedId, present := c.Params.Get("feedId")

	if !present {
		sendErrorResponse(c, 400, "Please provide Id of feed to unfollow")
	}

	feedUuid, parseErr := uuid.Parse(feedId)

	if parseErr != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Could not parse feedId: %v", parseErr))
	}

	err := apiCfg.DB.DeleteFeedsFollowedByUserByFeedId(c, database.DeleteFeedsFollowedByUserByFeedIdParams{
		UserID: user.ID,
		FeedID: feedUuid,
	})

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Couldn't delete feed: %v", err))
	}

	sendJsonResponse(c, 200, gin.H{
		"message": "Feed unfollowed successfully",
	})
}
