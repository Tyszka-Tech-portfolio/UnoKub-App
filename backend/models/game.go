package models

import (
	"errors" // Imported errors package
	"time"
)

type Game struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsFinished  bool      `json:"is_finished"`
	UserID      string    `json:"user_id"` // Added UserID field
}

func NewGame(id, name, description string) *Game {
	return &Game{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsFinished:  false,
	}
}

// ContinueGame fetches all unfinished games for a verified user
func ContinueGame(userID string, allGames []Game) ([]Game, error) {
	var unfinishedGames []Game
	for _, game := range allGames {
		if game.UserID == userID && !game.IsFinished {
			unfinishedGames = append(unfinishedGames, game)
		}
	}
	if len(unfinishedGames) == 0 {
		return nil, errors.New("No unfinished games found for this user")
	}
	return unfinishedGames, nil
}
