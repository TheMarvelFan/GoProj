package main

import (
	"fmt"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUserHandler(c *gin.Context) {
	type fields struct {
		Name string `json:"name"`
	}

	params := fields{}

	err := c.ShouldBindJSON(&params)

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error binding json: %v", err))
		return
	}

	user, createErr := apiCfg.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if createErr != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error creating user: %v", createErr))
		return
	}

	sendJsonResponse(c, 201, dbUserToUser(user))
}

func (apiCfg *apiConfig) getUserHandler(c *gin.Context, user database.User) {
	sendJsonResponse(c, 200, dbUserToUser(user))
}

func (apiCfg *apiConfig) getPostsForUserHandler(c *gin.Context, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(c, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		sendErrorResponse(c, 400, fmt.Sprintf("Error fetching posts for user: %v", err))
		return
	}

	sendJsonResponse(c, 200, dbPostsToPosts(posts))
}
