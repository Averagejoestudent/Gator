package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Averagejoestudent/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command,user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	info, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Could not run follow command : %w", err)
	}
	fmt.Printf("* Username:      %s\n", info.UserName)
	fmt.Printf("* Feedname:      %s\n", info.FeedName)
	return nil
}


func handlerFollowing(s *state, cmd command,user database.User) error {

	items , err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err != nil {
		return fmt.Errorf("couldn't find feed for current user: %w", err)
	}
	for _ , item := range items{
		fmt.Printf("* Feedname:      %s\n",item.FeedName)
	}
	return nil
}

func handlerUnFollow(s *state, cmd command,user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Could not unfollow because :%w", err)
	}
	return nil

}