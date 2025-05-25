package cli

import (
	"fmt"
	"context"
	"time"
	"encoding/json"
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/jayrgarg/gator/internal/state"
	"github.com/jayrgarg/gator/internal/database"
	"github.com/jayrgarg/gator/internal/rss"
)

func HandleAgg(s *state.State, cmd Command) error {

    if len(cmd.Args) != 1 {
        return fmt.Errorf("Expected 1 argument, got: %v", len(cmd.Args))
    }

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error parsing string time Duration, %v\n", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state.State) error {

	feedRow, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting next feed to fetch from DB, %v\n", err)
	}

	err = s.Db.MarkFeedFetched(context.Background(), feedRow.ID)
	if err != nil {
		return fmt.Errorf("Error marking feed as fetched from DB, %v\n", err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), feedRow.Url)

	for _, item := range rssFeed.Channel.Item {
		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("Error parsing time from rss feed item, %v\n", err)
			continue
		}
		utcParsedTime := parsedTime.In(time.UTC)
		var desc sql.NullString
		if item.Description != "" {
			desc = sql.NullString{String: item.Description, Valid: true}
		} else {
			desc = sql.NullString{String: "", Valid: false}
		}
		postParams := database.CreatePostParams{
			ID: 				uuid.New(),
			CreatedAt: 			time.Now(),
			UpdatedAt: 			time.Now(),
			Title: 				item.Title,
			Url: 				item.Link,
			Description: 		desc,
			PublishedAt: 		utcParsedTime,
			FeedID: 			feedRow.ID,
		}
		post, err := s.Db.CreatePost(context.Background(), postParams)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				// pqErr is now a *pq.Error, and you can access pqErr.Code
				if pqErr.Code == "23505" {
					fmt.Printf("Url already exists, %v\n", err)
				} else {
					fmt.Printf("Error in creating Post in DB, %v\n", err)
				}
			} else {
				fmt.Printf("Error in creating Post in DB, %v\n", err)
			}
			continue
		}
		jpost, err := json.MarshalIndent(post, "", "\t")
		if err != nil {
			return fmt.Errorf("Error marshall indenting post, %v\n", err)
		}
		fmt.Println("Creating Post: ")
		fmt.Println(string(jpost))

	}

	
	
	// jRssFeed, err := json.MarshalIndent(rssFeed, "", "\t")
	// if err != nil {
	// 	return fmt.Errorf("Error marshall indenting feed, %v\n", err)
	// }
	// fmt.Println(string(jRssFeed))

	return nil;
}
