package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/Nukie90/rssagg/internal/database"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Name     string   `json:"name"`
}

func databaseUsertoUser (dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name: dbUser.Name,
	}
}
