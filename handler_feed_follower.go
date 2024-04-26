package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/Nukie90/rssagg/internal/auth"
	"github.com/Nukie90/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (a *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}
	feedfollow, err := a.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:       uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		UserID:   user.ID,
		FeedID:   params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowtoFeedFollow(feedfollow))
}

func (a *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedfollows, err := a.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting feedfollows: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowstoFeedFollows(feedfollows))
}

func (a *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid feedFollowID: %v", err))
		return
	}

	err = a.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting feedfollow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}