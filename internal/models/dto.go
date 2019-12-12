package models

import (
	"time"
)

// State of request counter
type State struct {
	Total      int       `json:"total_count"`
	LastUpdate time.Time `json:"last_update"`
}
