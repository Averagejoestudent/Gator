package main

import (
	"context"
	"fmt"
	"time"
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
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	for i, item := range ptr_feed.Channel.Item {
		fmt.Printf("%d %s\n", i, item.Title)
	}

	return nil
}
