package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
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
		log.Println("New post found:", item.Title, "in feed:", feed.Name)
	}

	log.Printf("Feed %s was fetch. Found %v new posts.\n", feed.Name, len(rssFeed.Channel.Item))
}
