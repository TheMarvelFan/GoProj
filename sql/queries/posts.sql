-- name: CreatePost :one
INSERT INTO posts(
    id,
    created_at,
    updated_at,
    published_at,
    title,
    description,
    url,
    feed_id
) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetPostsForUser :many
SELECT posts.* FROM 
posts JOIN feed_followers ON
feed_followers.feed_id = posts.feed_id
WHERE feed_followers.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;