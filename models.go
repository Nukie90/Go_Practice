package main

import (
	"time"

	"github.com/Nukie90/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Name     string    `json:"name"`
	APIKey   string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name:     dbUser.Name,
		APIKey:   dbUser.ApiKey,
	}
}

type Feed struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	UserID   uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:       dbFeed.ID,
		CreateAt: dbFeed.CreateAt,
		UpdateAt: dbFeed.UpdateAt,
		Name:     dbFeed.Name,
		URL:      dbFeed.Url,
		UserID:   dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, feed := range dbFeeds {
		feeds[i] = databaseFeedtoFeed(feed)
	}
	return feeds
}