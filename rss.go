package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)

	if err != nil {
		return RSSFeed{}, err
	}

	defer res.Body.Close()

	data, errData := io.ReadAll(res.Body)

	if errData != nil {
		return RSSFeed{}, errData
	}

	rssFeed := RSSFeed{}

	errMarshal := xml.Unmarshal(data, &rssFeed)

	if errMarshal != nil {
		return RSSFeed{}, errMarshal
	}

	return rssFeed, nil
}
