package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nukie90/rssagg/internal/database"
	_"github.com/Nukie90/rssagg/internal/auth"
	"github.com/google/uuid"
)

func (a *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:  uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name: params.Name,
		
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (a *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

