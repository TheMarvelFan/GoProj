-- name: CreateFeedFollowers :one
INSERT INTO feed_followers(id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedsFollowedByUser :many
SELECT * FROM feed_followers WHERE user_id = $1;

-- name: GetFollowersForFeed :many
SELECT * FROM feed_followers WHERE feed_id = $1;