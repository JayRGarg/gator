-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
WITH user_feed_follows AS (
    SELECT feed_id
    FROM feed_follows
    WHERE feed_follows.user_id = $1
) 
SELECT posts.*
FROM posts
INNER JOIN user_feed_follows
ON posts.feed_id = user_feed_follows.feed_id
ORDER BY posts.published_at DESC
LIMIT $2;
