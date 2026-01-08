package main

import (
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

type FeedFollower struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FeedID    uuid.UUID `json:"feedId"`
	UserID    uuid.UUID `json:"userId"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func dbFeedFollowerToFeedFollower(dbFeedFollower database.FeedFollower) FeedFollower {
	return FeedFollower{
		ID:        dbFeedFollower.ID,
		CreatedAt: dbFeedFollower.CreatedAt,
		UpdatedAt: dbFeedFollower.UpdatedAt,
		FeedID:    dbFeedFollower.FeedID,
		UserID:    dbFeedFollower.UserID,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(feed))
	}

	return feeds
}

func dbUsersToUsers(dbUsers []database.User) []User {
	users := []User{}

	for _, user := range dbUsers {
		users = append(users, dbUserToUser(user))
	}

	return users
}
