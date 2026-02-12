-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;



-- name: UsernameFeed :many
SELECT users.name as username , feeds.name , feeds.url
from feeds
INNER join users
ON feeds.user_id = users.id; 

-- name: GetFeedByURL :one
Select * FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1;


-- name: GetNextFeedToFetch :one
Select * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
