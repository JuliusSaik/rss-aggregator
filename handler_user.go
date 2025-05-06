package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JuliusSaik/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error prasing JSON", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not create user", err, user))
	}

	respondWithJSON(w, 201, databaseToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, 200, databaseToUser(user))
}
