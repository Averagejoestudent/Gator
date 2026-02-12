package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Averagejoestudent/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <Time to request>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %s...\n", cmd.Args[0])
	for ; ; <-ticker.C {
		err := scrapefeeds(s)
		if err != nil {
			fmt.Printf("There is error in scraping %s\n", err)
		}
	}
	return nil
}

func scrapefeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Cannot scrape feed")
	}
	s.db.MarkFeedFetched(context.Background(), feed.ID)
	ptr_feed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("couldn't create post: %s\n", err)
	}

	for _, item := range ptr_feed.Channel.Item {
		parsedTime, parseErr := time.Parse(time.RFC1123, item.PubDate)
		if parseErr != nil {
			fmt.Printf("Publication date convertion issue: %s\n", parseErr)
		}
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: parsedTime, Valid: parseErr == nil},
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			fmt.Printf("couldn't create post: %s\n", err)
		}

	}

	return nil
}
