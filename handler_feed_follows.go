package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JuliusSaik/rss-aggregator/internal/db"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error prasing JSON", err))
		return
	}

	feed_follows, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not create feed follow", err))
		return
	}

	respondWithJSON(w, 201, feed_follows)
}

func (apiCfg *apiConfig) handlerGetFeedFollowsByUserId(w http.ResponseWriter, r *http.Request, user db.User) {
	feed_follows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not get feed follows", err))
		return
	}

	respondWithJSON(w, 200, feed_follows)
}

func (apiCfg *apiConfig) handlerDeleteFeeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	feedUUID, err := uuid.Parse(feedFollowId)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not parse feed follow Id", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedUUID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not delete feed follow", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
