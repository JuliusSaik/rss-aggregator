package main

import (
	"database/sql"
	"time"

	"github.com/JuliusSaik/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID      `json:"id"`
	Username     sql.NullString `json:"username"`
	PasswordHash sql.NullString `json:"password_hash"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Name         string         `json:"name"`
}

func databaseToUser(dbUser db.User) User {
	return User{
		ID:           dbUser.ID,
		Username:     dbUser.Username,
		PasswordHash: dbUser.PasswordHash,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		Name:         dbUser.Name,
	}
}
