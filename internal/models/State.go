package models

import (
	"time"
)

// State of counter
type State struct {
	Total      int       `json:"total_count"`
	LastUpdate time.Time `json:"last_update"`
}
