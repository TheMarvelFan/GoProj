package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	scraperThreadsCount int,
	requestDelay time.Duration,
) {
	log.Printf("Scraping on %v goroutines with %s delay between requests\n", scraperThreadsCount, requestDelay.String())
	ticker := time.NewTicker(requestDelay)

	for ; ; <-ticker.C { // if we used for range ticker.C then it would block the main goroutine
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(scraperThreadsCount))

		if err != nil {
			log.Println("Error fetching feeds:", err)
			continue
		}

		wg := sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, &wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(
	db *database.Queries,
	wg *sync.WaitGroup,
	feed database.Feed,
) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Println("Error marking feed fetched:", err)
		return
	}

	rssFeed, feedConvErr := urlToFeed(feed.Url)

	if feedConvErr != nil {
		log.Println("Error fetching feed:", feedConvErr)
	}

	for _, item := range rssFeed.Channel.Item {
		nullableDescription := sql.NullString{}

		if item.Description != "" {
			nullableDescription.String = item.Description
			nullableDescription.Valid = true
		}

		publishedAt, timeConvError := time.Parse(time.RFC1123Z, item.PubDate)

		if timeConvError != nil {
			log.Printf("Cannot parse published date/time %v , error: %v", item.PubDate, timeConvError)
			continue
		}

		_, createPostErr := db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				PublishedAt: publishedAt,
				Title:       item.Title,
				Description: nullableDescription,
				Url:         item.Link,
				FeedID:      feed.ID,
			},
		)

		if createPostErr != nil {
			if strings.Contains(createPostErr.Error(), "duplicate key") { // this is where we should ideally set log levels
				continue
			}

			log.Println("Error creating post:", createPostErr)
		}
	}

	log.Printf("Feed %s was fetch. Found %v new posts.\n", feed.Name, len(rssFeed.Channel.Item))
}
