package services

import "github.com/bgoldovsky/counter-api/internal/models"

// Counter service
type Counter interface {
	Increment() error
	GetState() models.State
}
