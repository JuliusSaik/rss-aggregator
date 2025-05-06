package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/JuliusSaik/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func startScraping(
	queriesDatabase *db.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scraping on %v gouroutines, every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := queriesDatabase.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error getting feeds to fetch", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(queriesDatabase, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(queriesDatabase *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	err := queriesDatabase.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed from url", feed.Url, "error:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Could not parse date", item.PubDate, err)
			return
		}

		_, err = queriesDatabase.CreatePost(context.Background(), db.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			FeedID:      feed.ID,
			Description: description,
			PublishedAt: parsedTime,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Could not create post")
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
