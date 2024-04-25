package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/Nukie90/rssagg/internal/auth"
	"github.com/Nukie90/rssagg/internal/database"
	"github.com/google/uuid"
)

func (a *apiConfig) handlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}
	feed, err := a.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name:     params.Name,
		Url:      params.URL,
		UserID:   user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedtoFeed(feed))
}

func (a *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := a.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}
