package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Averagejoestudent/Gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var err error
	limit := 2
	if len(cmd.Args) > 0 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("cannot convert argument string into integer limit :%w", err)
		}
	}
	Posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}
	for _, Post := range Posts {
		fmt.Printf("* Title:      %s\n", Post.Title)
		fmt.Println("---")
		fmt.Printf("* Description:      %s\n", Post.Description.String)
		fmt.Println("---")
		if Post.PublishedAt.Valid {
			fmt.Printf("* Published:	%s\n", Post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("* Published:	%s\n", "N/A")
		}
		fmt.Printf("* URL: %s\n", Post.Url)
		fmt.Println("---")
	}

	return nil
}
